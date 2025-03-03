package main

import (
	"example.com/api/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Run(":8080")
	routers.RegisterRoutes(server)
}
