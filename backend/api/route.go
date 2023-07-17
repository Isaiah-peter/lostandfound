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
	authRouter.PUT("/category/update/:id", server.updateCategory)
	authRouter.DELETE("/category/delete/:id", server.deleteCategory)
	authRouter.POST("/lostitems", server.createLostItem)
	authRouter.GET("/lostitems/:id", server.getLostItem)
	authRouter.GET("/lostitems/list", server.ListLostItem)
	authRouter.PUT("/lostitems/update/:id", server.updateLostItem)
	authRouter.DELETE("/lostitems/delete/:id", server.deleteLostItem)

	server.router = router
}