package main

import (
	"context"
	"net"
	"net/http"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	db "github.com/CineDeepMatch/Backend-server/db/sqlc"
	util "github.com/CineDeepMatch/Backend-server/db/utils"
	"github.com/CineDeepMatch/Backend-server/gapi"
	"github.com/CineDeepMatch/Backend-server/mail"
	mongodb "github.com/CineDeepMatch/Backend-server/mongodb/repositories"
	"github.com/CineDeepMatch/Backend-server/pb"
	"github.com/CineDeepMatch/Backend-server/worker"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")

	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	connMongoDB, err := mongo.Connect(context.TODO(), options.Client().ApplyURI((config.MongoDBSource)))

	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to MongoDB")
	}

	store := db.NewStore(connPool)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisServerAddress,
	}
	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	mongoDBStore := mongodb.NewStore(connMongoDB, "CineDeepMatch", "Movies")

	go runTaskProcessor(config, redisOpt, store)
	go runGatewayServer(config, store, mongoDBStore, taskDistributor)
	runGrpcServer(config, store, mongoDBStore, taskDistributor)

}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")

	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")

	}
	log.Info().Msg("db migrated successfully")
}

func runTaskProcessor(config util.Config, redisOpt asynq.RedisClientOpt, store db.Store) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)
	log.Info().Msg("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}

func runGrpcServer(config util.Config, store db.Store, mongoDBStore mongodb.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, store, mongoDBStore, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCineDeepMatchServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}

func runGatewayServer(config util.Config, store db.Store, mongoDBStore mongodb.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, store, mongoDBStore, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterCineDeepMatchHandlerServer(ctx, grpcMux, server)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot register handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	fs := http.FileServer(http.Dir("./doc/swagger"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	handler := cors.Default().Handler(mux)
	log.Info().Msgf("start gRPC server at %s", config.HTTPServerAddress)

	err = http.ListenAndServe(config.HTTPServerAddress, handler)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot start HTTP gateway server")
	}
}

// func runGrpcServerWithAuth(config util.Config, store db.Store, mongoDBStore mongodb.Store) {
// 	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

// 	if err != nil {
// 		log.Fatal().Err(err).Msg("cannot create token maker")
// 	}

// 	server, err := gapi.NewAuthServer(config, store, mongoDBStore, tokenMaker)
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("cannot create server")
// 	}

// 	interceptor := gapi.NewAuthInterceptor(tokenMaker)

// 	serverOptions := []grpc.ServerOption{
// 		grpc.UnaryInterceptor(interceptor.Unary()),
// 	}

// 	grpcServer := grpc.NewServer(serverOptions...)

// 	pb.RegisterCineDeepMatchServer(grpcServer, server)
// 	reflection.Register(grpcServer)

// 	listener, err := net.Listen("tcp", config.GRPCServerAddress)

// 	if err != nil {
// 		log.Fatal().Err(err).Msg("cannot create listener")
// 	}

// 	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())

// 	err = grpcServer.Serve(listener)

// 	if err != nil {
// 		log.Fatal().Err(err).Msg("cannot start gRPC server")
// 	}
// }
