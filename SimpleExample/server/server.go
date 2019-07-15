package server

import (
	"fmt"

	"github.com/fvbock/endless"
)

// Start ...
func Start() {
	if err := endless.ListenAndServe("0.0.0.0:8888", SetupRouter()); err != nil {
		fmt.Printf("Failed to listen and serve: %s", err)
	}
}
