package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, err := url.Parse("http://localhost:8080")
	if err != nil {
		log.Fatal("Failed to parse target url")
}
	proxy := httputil.NewSingleHostReverseProxy(target)

	director := proxy.Director
	proxy.Director = func(r *http.Request) {
		director(r)
		log.Printf("Proxing %s %s -> %s", r.Method, r.URL.Path, target.Host)
	}

	http.Handle("/", proxy)

	log.Print("Starting ML proxi on port 8000")
	http.ListenAndServe(":8000", nil)
}
