package utility

import (
	"AdvanceExampleWithES/modules/request"
	"bytes"
	"fmt"
)

// ForwardRequest is responsible to decide where request will be send.
// Request is send to Elasticsearch and RabbitMQ by default.
func ForwardRequest(data map[interface{}]interface{}) {
	ECData := ToECFormat(data)
	sendToElasticsearch(ECData)
}

func sendToElasticsearch(data string) {
	if err := request.Request("PUT", "http://127.0.0.1:9200/_bulk", bytes.NewBuffer([]byte(data))); err != nil {
		fmt.Println(err)
	}
}
