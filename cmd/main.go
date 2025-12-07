package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"go-ml-router/pkg/config"
)

func main() {
	path := "router_config.yaml"

	config, err := config.FromYaml(path)
	if err != nil {
		log.Fatalf("Failed to read config")
	}

	if err != nil {
		log.Fatal("Failed to parse target url")
}
	target := config.PrimaryBackend
	proxy := httputil.NewSingleHostReverseProxy(target.Url())

	director := proxy.Director
	proxy.Director = func(r *http.Request) {
		director(r)
		log.Printf("Proxing %s %s -> %s", r.Method, r.URL.Path, target.Address)
	}

	http.Handle("/", proxy)

	log.Print("Starting ML proxi on port 8000")
	http.ListenAndServe(":8000", nil)
}
