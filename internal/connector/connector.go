package connector

import (
	"net/http"
	"time"
)

// NewConnection initialises an HTTP client with proxy support from environment.
func NewConnection() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		Proxy:              http.ProxyFromEnvironment,
	}
	return &http.Client{Transport: tr, Timeout: 10 * time.Second}
}
