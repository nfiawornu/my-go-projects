package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/nfiawornu/my-go-projects/simple_bank/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// Accounts api's
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	//Transfer api's
	router.POST("/transfers", server.createTransfer)
	router.GET("/transfers/:id", server.getTransfer)
	router.GET("transfers", server.listTransfers)

	//Entry api's
	router.POST("entries", server.createEntry)
	router.GET("/entries/:id", server.getEntry)
	router.GET("entries", server.listEntries)

	server.router = router
	return server
}

// start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
