package utility

import (
	"SimpleExampleWithES/modules/request"
	"bytes"
	"fmt"
	"time"
)

// SendToElasticsearch ...
func SendToElasticsearch(date time.Time, data string) {
	fmt.Println(data)
	if err := request.Request("PUT", fmt.Sprintf("http://127.0.0.1:9200/s3/log2/%v", date), bytes.NewBuffer([]byte(data))); err != nil {
		fmt.Println(err)
	}
}
