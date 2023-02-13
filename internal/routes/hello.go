package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHelloWorldRoutes registers the routes for the hello world
func HelloWorld(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Hello World 2!"})
}
