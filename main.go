package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Run(":8080")
	fmt.Println("test")
}
