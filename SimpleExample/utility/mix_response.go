package utility

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// MixProxyResponse ...
func MixProxyResponse(c *gin.Context, resp *http.Response) {
	for idx, val := range resp.Header {
		c.Writer.Header().Set(idx, strings.Join(val, " "))
		if idx == "Etag" {
			c.Writer.Header().Del("Etag")
			c.Writer.Header()["ETag"] = val
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(0)
	}
	defer resp.Body.Close()

	c.Status(resp.StatusCode)

	c.Writer.Write(body)
}
