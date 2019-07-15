package main

import (
	"AdvanceExampleWithES/modules/cache"
	"AdvanceExampleWithES/modules/interval"
	"AdvanceExampleWithES/server"
	"AdvanceExampleWithES/utility"
	"fmt"
)

func main() {
	cache.New("ESCache", 10).Build()
	interval.Create(5, func() {
		if cData := cache.Use("ESCache").GetALL(true); len(cData) != 0 {
			utility.ForwardRequest(cData)
			fmt.Println("Updated Elasticsearch!")
		}
	})
	server.Start()
}
