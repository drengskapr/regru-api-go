package connector

import (
	"net/http"
	"testing"
	"time"
)

func TestNewConnection(t *testing.T) {
	c := NewConnection()

	if c == nil {
		t.Fatal("expected non-nil *http.Client")
	}
	if c.Timeout != 10*time.Second {
		t.Errorf("expected timeout 10s, got %v", c.Timeout)
	}
	tr, ok := c.Transport.(*http.Transport)
	if !ok {
		t.Fatal("expected *http.Transport")
	}
	if tr.MaxIdleConns != 10 {
		t.Errorf("expected MaxIdleConns 10, got %d", tr.MaxIdleConns)
	}
	if tr.IdleConnTimeout != 30*time.Second {
		t.Errorf("expected IdleConnTimeout 30s, got %v", tr.IdleConnTimeout)
	}
	if !tr.DisableCompression {
		t.Error("expected DisableCompression true")
	}
	if tr.Proxy == nil {
		t.Error("expected Proxy to be set (http.ProxyFromEnvironment)")
	}
}
