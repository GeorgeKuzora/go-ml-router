package server

import (
	"go-ml-router/pkg/config"
	"go-ml-router/pkg/proxy"
	"log"
	"net/http"
	"time"
)

func Serve(config *config.App, proxyManager *proxy.ProxyManager) {

	mux := http.NewServeMux()
	mux.Handle("/predict", proxyManager.ProxyHandler())

    // Start server
    server := &http.Server{
        Addr:         ":8080",
        Handler:      mux,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 30 * time.Second,
        IdleTimeout:  120 * time.Second,
    }

    log.Printf("Starting ML Proxy on %s", server.Addr)
    err := server.ListenAndServe()
    log.Fatal(err)
}
