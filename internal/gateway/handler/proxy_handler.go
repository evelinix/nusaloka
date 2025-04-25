package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// NewReverseProxy membuat handler proxy untuk service tertentu
func NewReverseProxy(target string) http.HandlerFunc {
	parsedURL, _ := url.Parse(target)

	// Inisialisasi proxy
	proxy := httputil.NewSingleHostReverseProxy(parsedURL)

	// Modifikasi request URL sebelum diteruskan
	proxy.Director = func(req *http.Request) {
		req.URL.Path = strings.TrimPrefix(req.URL.Path, "/api/account")
		req.Host = parsedURL.Host
	}

	return proxy.ServeHTTP
}
