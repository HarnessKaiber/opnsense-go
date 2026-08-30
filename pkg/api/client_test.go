package api

import (
	"net/http"
	"testing"
)

func TestNewClientEnablesHTTP2(t *testing.T) {
	client := NewClient(Options{})
	transport, ok := client.client.HTTPClient.Transport.(*http.Transport)
	if !ok {
		t.Fatalf("transport type = %T, want *http.Transport", client.client.HTTPClient.Transport)
	}
	if !transport.ForceAttemptHTTP2 {
		t.Fatal("ForceAttemptHTTP2 is disabled")
	}
}
