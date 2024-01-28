package gapi

import (
	db "github.com/CineDeepMatch/Backend-server/db/sqlc"
	util "github.com/CineDeepMatch/Backend-server/db/utils"
	mongodb "github.com/CineDeepMatch/Backend-server/mongodb/repositories"
	"github.com/CineDeepMatch/Backend-server/pb"
	"github.com/CineDeepMatch/Backend-server/token"
)

type AuthServer struct {
	pb.UnimplementedCineDeepMatchServer
	config       util.Config
	store        db.Store
	tokenMaker   token.Maker
	mongoDBStore mongodb.Store
}

func NewAuthServer(config util.Config, store db.Store, mongoDBStore mongodb.Store, tokenMaker token.Maker) (*AuthServer, error) {
	server := &AuthServer{
		config:       config,
		store:        store,
		mongoDBStore: mongoDBStore,
		tokenMaker:   tokenMaker,
	}

	return server, nil
}
