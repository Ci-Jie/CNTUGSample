package server

import (
	"SimpleExample/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ...
func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.NoRoute(controllers.S3TempFunction)

	return router
}
