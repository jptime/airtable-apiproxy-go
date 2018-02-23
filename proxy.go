package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// Necessary for https
type Transport struct{}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	apiKey := os.Getenv("API_KEY")
	req.Host = req.URL.Host
	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Header.Set("Access-Control-Allow-Headers", "X-Requested-With")
	req.Header.Add("Authorization", "Bearer "+apiKey)

	return http.DefaultTransport.RoundTrip(req)
}

// Sets the path to the table to fetch
func NewReverseProxy(target string, path string) *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "https",
		Host:   target,
		Path:   path,
	})
}

func Handle(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")

		p.ServeHTTP(w, r)
	}
}
