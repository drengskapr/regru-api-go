package client

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestApiRequest_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result":"success"}`))
	}))
	defer srv.Close()

	body, err := ApiRequest(srv.URL, map[string]string{"key": "val"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(body) != `{"result":"success"}` {
		t.Errorf("unexpected body: %s", body)
	}
}

func TestApiRequest_HTTPError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`server error`))
	}))
	defer srv.Close()

	_, err := ApiRequest(srv.URL, map[string]string{})
	if err == nil {
		t.Fatal("expected error for 5xx response")
	}
	if !strings.Contains(err.Error(), "500") {
		t.Errorf("expected error to mention status 500, got: %v", err)
	}
}

func TestApiRequest_ConnectionError(t *testing.T) {
	_, err := ApiRequest("http://127.0.0.1:0/unreachable", map[string]string{})
	if err == nil {
		t.Fatal("expected error for unreachable server")
	}
}
