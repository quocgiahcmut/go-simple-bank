package gapi

import (
	"fmt"

	db "github.com/quocgiahcmut/simple-bank/db/sqlc"
	"github.com/quocgiahcmut/simple-bank/pb"
	"github.com/quocgiahcmut/simple-bank/token"
	"github.com/quocgiahcmut/simple-bank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// New Server creates a new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, err
}
