package main

import (
	"fmt"

	"github.com/outofforest/run"
)

func main() {
	run.Service("cache", nil, func() error {
		fmt.Println("test")
		return nil
	})
}
