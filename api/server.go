package api

import (
	db "backend-master-class/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.POST("/accounts", server.createAccount)
	router.PUT("/accounts/:id", server.updateAccount)
	router.DELETE("accounts/:id", server.deleteAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(adress string) error {
	return server.router.Run(adress)
}

// errorResponse return a json error response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
