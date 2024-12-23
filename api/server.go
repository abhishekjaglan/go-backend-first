package api

import (
	db "github.com/abhishekjaglan/go-backend-first/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

// Starts the http server on specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// formatting error response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
