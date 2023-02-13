package main

import (
	"app/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", routes.HelloWorld)

	router.Run(":8000")
}
