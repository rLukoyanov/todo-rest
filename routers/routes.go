package routers

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/list", getTasks)
	server.POST("/create", createTask)
	server.DELETE("/delete", deleteTask)
	server.PUT("/done", updateTask)
}
