package api

import (
	"fmt"

	db "github.com/Isaiah-peter/lostandfound/db/sqlc"
	"github.com/Isaiah-peter/lostandfound/token"
	"github.com/gin-gonic/gin"
)



type Server struct {
	store *db.Store
	tokenMaker token.Maker
	router *gin.Engine
}

func NewServer(store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetomaker("qwertyqwerty")
	if err != nil {
		return nil, fmt.Errorf("cannot make token %w", err)
	}
	server := &Server{
		store: store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	router.POST("/user", server.createUser)
	router.GET("/user/:id", server.getUser)

	server.router = router
	return server, nil
}

func (server *Server) Start(adress string) error {
	return server.router.Run(adress)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}