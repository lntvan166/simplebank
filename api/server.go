package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	db "github.com/lntvan166/simplebank/db/sqlc"
)

type Server struct {
	store db.Store
	route *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	route := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	route.POST("/accounts", server.createAccount)
	route.GET("/accounts/:id", server.getAccount)
	route.GET("/accounts", server.listAccount)

	route.POST("/transfers", server.createTransfer)

	server.route = route

	return server
}

func (server *Server) Start(address string) error {
	return server.route.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
