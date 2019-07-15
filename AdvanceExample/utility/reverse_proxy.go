package utility

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// ReverseProxy ...
func ReverseProxy(c *gin.Context) {
	url, _ := url.Parse("http://172.17.1.100:7480")

	director := func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
	}

	proxy := &httputil.ReverseProxy{Director: director}

	defer c.Request.Body.Close()

	proxy.ServeHTTP(c.Writer, c.Request)
}
