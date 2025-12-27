package server

import (
	"fmt"
	"go-ml-router/pkg/config"
	"go-ml-router/pkg/proxy"
	"log"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(proxyManager *proxy.ProxyManager, config *config.App) Server {
	mux := http.NewServeMux()
	mux.Handle(config.Routes.Predict, proxyManager.ProxyHandler())

	server := &http.Server{
	    Addr:         fmt.Sprintf(":%d", config.Port),
	    Handler:      mux,
	    ReadTimeout:  config.ReadTimeout(),
	    WriteTimeout: config.WriteTimeout(),
	    IdleTimeout:  config.IdleTimeout(),
	}

	return Server{server: server}
}

func (s *Server) Serve() error {
    log.Printf("Starting ML Proxy on %s", s.server.Addr)
    err := s.server.ListenAndServe()
    return err
}
