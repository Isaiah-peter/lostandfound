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
	tokenMaker, err := token.NewPasetomaker("qawsedrftgyhujikolp1z2x3c4v5b6n7")
	if err != nil {
		return nil, fmt.Errorf("cannot make token %w", err)
	}
	server := &Server{
		store: store,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}


func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}