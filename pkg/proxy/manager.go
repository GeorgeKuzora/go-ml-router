package proxy

import (
	"net/http"
	"net/http/httputil"
	"sync"
	"time"
)

type ProxyManager struct {
	proxies map[string]*httputil.ReverseProxy
	mu sync.RWMutex
	router *Router
}

func NewProxyManager(router *Router) ProxyManager {
	return ProxyManager{
		proxies: make(map[string]*httputil.ReverseProxy),
		router: router,
	}
}

func (pm *ProxyManager) GetProxy(backendKey string) (*httputil.ReverseProxy, error) {
	pm.mu.RLock()
	proxy, exists := pm.proxies[backendKey]
	pm.mu.Unlock()

	if exists {
		return proxy, nil
	}

	pm.mu.Lock()
	defer pm.mu.Unlock()

	if proxy, exists = pm.proxies[backendKey]; exists {
		return proxy, nil
	}

	backendUrl, err := pm.router.GetBackendUrl(backendKey)
	if err != nil {
		return nil, err
	}
	proxy = httputil.NewSingleHostReverseProxy(&backendUrl)

	proxy.Transport = &http.Transport{
		MaxIdleConns: 100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout: 90 * time.Second,
	}

	originalDirector := proxy.Director

	proxy.Director = func(r *http.Request) {
		originalDirector(r)
		r.Header.Set("X-Proxy", "ml-proxy-v1")
		r.Header.Set("X-Backend-Key", backendKey)
	}

	pm.proxies[backendKey] = proxy

	return proxy, nil
}

func (pm *ProxyManager) ProxyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		model := r.Header.Get("X-Model-Name")
		if model == "" {
			model = "default"
		}

		backendKey, err := pm.router.SelectBackend(model)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		proxy, err := pm.GetProxy(backendKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		proxy.ServeHTTP(w, r)
	}
}
