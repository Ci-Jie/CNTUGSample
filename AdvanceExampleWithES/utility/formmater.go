package utility

import (
	"AdvanceExampleWithES/modules/cache"
	"bytes"
	"encoding/json"
	"strings"
)

// ECMeta is defined metadata about Elasticsearch.
type ECMeta struct {
	Index ECMetaIndex `json:"index"`
}

// ECMetaIndex is a part of metadata in ECMeta.
type ECMetaIndex struct {
	Index string      `json:"_index"`
	Type  string      `json:"_type"`
	ID    interface{} `json:"_id"`
}

var dataArray []string

// ToECFormat is responsible to transfer format for Elasticsearch metadata.
func ToECFormat(data map[interface{}]interface{}) string {
	for key, item := range data {
		metadata, _ := json.Marshal(ECMeta{Index: ECMetaIndex{Index: "s3", Type: "log", ID: key}})
		dataArray = append(dataArray, bytes.NewBuffer(metadata).String())
		data, _ := json.Marshal(item)
		dataArray = append(dataArray, bytes.NewBuffer(data).String())

		cache.Use("ESCache").Remove(key)
	}
	resp := strings.Join(dataArray, "\n") + "\n"
	dataArray = nil
	return resp
}
