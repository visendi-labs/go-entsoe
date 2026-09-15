package entsoe

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRequestTimesOutRatherThanHanging(t *testing.T) {
	blocked := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-blocked
	}))
	// Unblock the handler before shutting the server down, or Close waits on it.
	defer srv.Close()
	defer close(blocked)

	c := NewEntsoeClient("token")
	c.SetHTTPClient(&http.Client{Timeout: 100 * time.Millisecond})

	done := make(chan error, 1)
	go func() {
		_, err := c.httpClient.Get(srv.URL)
		done <- err
	}()

	select {
	case err := <-done:
		if err == nil {
			t.Fatal("expected a timeout error, got nil")
		}
	case <-time.After(5 * time.Second):
		t.Fatal("request hung instead of timing out")
	}
}

func TestDefaultClientHasATimeout(t *testing.T) {
	if got := NewEntsoeClient("token").httpClient.Timeout; got != DefaultTimeout {
		t.Fatalf("default client timeout = %v, want %v", got, DefaultTimeout)
	}
}
