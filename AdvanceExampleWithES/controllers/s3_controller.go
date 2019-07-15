package controllers

import (
	"AdvanceExampleWithES/modules/cache"
	"AdvanceExampleWithES/utility"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type baseCache struct {
	Date   time.Time `json:"date"`
	Method string    `json:"method"`
	Status int       `json:"status"`
	URL    string    `json:"url"`
	Bucket string    `json:"bucket"`
}

// ReverseProxyFunction ...
func ReverseProxyFunction(c *gin.Context) {
	utility.ReverseProxy(c)

	id := time.Now().Format("2006-01-02T15:04:05.999Z")

	date, err := time.Parse("2006-01-02T15:04:05.999Z", id)

	if err != nil {
		fmt.Println(err)
		return
	}

	method := c.Request.Method
	status := c.Writer.Status()
	url := fmt.Sprintf("%v", c.Request.URL)
	bucketName := getBucketName(url)

	if err := cache.Use("ESCache").Set(id, baseCache{date, method, status, url, bucketName}); err != nil {
		fmt.Println(err)
	}
}

// getBucketName is responsible to return bucket name from request url.
func getBucketName(url string) string {
	bucketName := strings.Split(url, "/")[1]
	return bucketName
}
