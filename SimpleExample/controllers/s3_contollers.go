package controllers

import (
	"SimpleExample/utility"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func S3TempFunction(c *gin.Context) {
	resp, err := utility.SendRequest(c.Request)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(0)
	}

	utility.MixProxyResponse(c, resp)
}
