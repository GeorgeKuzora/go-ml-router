package main

import (
	"go-ml-router/pkg/config"
	"go-ml-router/pkg/proxy"
	"go-ml-router/pkg/server"
	"log"
)

func main() {
	path := "config.yaml"

	config, err := config.FromYaml(path)
	if err != nil {
		log.Fatalf("Failed to read config")
	}

	router := proxy.NewRouter(config)
	if err != nil {
		log.Fatalf("Failed to start a router")
	}

	proxyManager := proxy.NewProxyManager(&router)

	server.Serve(&config.App, &proxyManager)
}
