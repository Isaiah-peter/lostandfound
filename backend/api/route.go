package api

import "github.com/gin-gonic/gin"

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)

	authRouter := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRouter.GET("/user/:id", server.getUser)
	authRouter.POST("/category", server.createCategory)
	authRouter.GET("/category/:id", server.getCategory)
	authRouter.GET("/category/list", server.ListCategory)

	server.router = router
}