package controllers

import (
	"AdvanceExample/utility"

	"github.com/gin-gonic/gin"
)

// ReverseProxyFunction ...
func ReverseProxyFunction(c *gin.Context) {
	utility.ReverseProxy(c)
}
