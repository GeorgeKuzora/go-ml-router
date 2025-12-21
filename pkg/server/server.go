package server

import (
	"fmt"
	"go-ml-router/pkg/config"
	"go-ml-router/pkg/proxy"
	"log"
	"net/http"
)

func Serve(config config.App, router proxy.Router) {
	http.Handle("/", router.Proxy)

	log.Print("Starting ML proxi on port 8000")
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
