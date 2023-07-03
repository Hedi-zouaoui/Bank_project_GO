package gapi

import (
	"fmt"

	db "github.com/Hedi-zouaoui/go_project/db/sqlc"
	"github.com/Hedi-zouaoui/go_project/pb"
	"github.com/Hedi-zouaoui/go_project/token"
	util "github.com/Hedi-zouaoui/go_project/utils"
)

type Server struct {
	pb.UnimplementedSimpleBankServer 
	config util.Config
	store  db.Store

	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker("12345678901234567890123456789012")
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker : %w ", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil

}
