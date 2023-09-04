package main

import (
	"fmt"

	"github.com/GuilhermeSa/go-service-template/internal/server"
)

func main() {
	config, err := server.LoadConfig("config.yaml")
	if err != nil {
		panic(fmt.Errorf("loading config: %w", err))
	}
	server.Start(config)
}
