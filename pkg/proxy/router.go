package proxy

import (
	"go-ml-router/pkg/config"
	"net/http/httputil"
	"net/url"
)

type Router struct {
	Backends map[string]*url.URL
	Proxy *httputil.ReverseProxy
}

func (r Router) GetBackendUrl(backendKey string) (url.URL, error) {

}

func (r Router) SelectBackend(model string) (string, error) {

}

func NewRouter(config config.Config) Router {
	return Router{}
}
