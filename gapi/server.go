package gapi

import (
	"fmt"

	db "github.com/CineDeepMatch/Backend-server/db/sqlc"
	util "github.com/CineDeepMatch/Backend-server/db/utils"
	"github.com/CineDeepMatch/Backend-server/pb"
	"github.com/CineDeepMatch/Backend-server/token"
)

type Server struct {
	pb.UnimplementedCineDeepMatchServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
