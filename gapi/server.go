package gapi

import (
	"fmt"

	db "github.com/CineDeepMatch/Backend-server/db/sqlc"
	util "github.com/CineDeepMatch/Backend-server/db/utils"
	mongodb "github.com/CineDeepMatch/Backend-server/mongodb/repositories"
	"github.com/CineDeepMatch/Backend-server/pb"
	"github.com/CineDeepMatch/Backend-server/token"
	"github.com/CineDeepMatch/Backend-server/worker"
)

type Server struct {
	pb.UnimplementedCineDeepMatchServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
	mongoDBStore    mongodb.Store
}

func NewServer(config util.Config, store db.Store, mongoDBStore mongodb.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		mongoDBStore:    mongoDBStore,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
