package server

import (
	"SimpleExampleWithES/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ...
func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.NoRoute(controllers.ReverseProxyFunction)

	return router
}
