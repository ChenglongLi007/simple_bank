package gapi

import (
	"fmt"

	db "github.com/simple_bank/db/sqlc"
	"github.com/simple_bank/pb"
	"github.com/simple_bank/token"
	"github.com/simple_bank/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

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

	return server, nil
}
