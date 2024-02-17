package web

import "github.com/gin-gonic/gin"

func RegisterRoutes() *gin.Engine {
	server := gin.Default()
	registerUsersRoutes(server)
	return server
}
func registerUsersRoutes(server *gin.Engine) {
	u := &UserHandle{}
	server.POST("/users/signup", u.SignUp)
	server.POST("/users/login", u.Login)
	server.POST("/users/edit", u.Edit)
	//server.GET("/users/:id", func(context *gin.Context) {
	//
	//})
	server.GET("/users/profile", u.Profile)
}
