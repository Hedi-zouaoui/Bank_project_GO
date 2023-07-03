package api

import (
	"fmt"

	db "github.com/Hedi-zouaoui/go_project/db/sqlc"
	"github.com/Hedi-zouaoui/go_project/token"
	util "github.com/Hedi-zouaoui/go_project/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
)

type Server struct {
	config     util.Config
	store      db.Store
	router     *gin.Engine
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

	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
router.POST("/tokens/renew_access" , server.renewAccessToken)
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/account", server.createAccount)
	authRoutes.GET("/account/:id", server.getaccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.POST("/transfers", server.createTransfer)
	server.router = router
	return server, nil

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
