package webexamples

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRequestIDMiddleware(t *testing.T) {
	handler := RequestIDMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if RequestID(r.Context()) == "" {
			t.Fatal("request id missing from context")
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	handler.ServeHTTP(recorder, request)
	if recorder.Header().Get("X-Request-ID") == "" {
		t.Fatal("request id missing from response")
	}
}

func TestNewServerTimeouts(t *testing.T) {
	server := NewServer(":0", http.NewServeMux(), ServerConfig{
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       2 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       4 * time.Second,
	})
	if server.ReadHeaderTimeout != time.Second || server.IdleTimeout != 4*time.Second {
		t.Fatalf("unexpected timeout config: %+v", server)
	}
}

func TestBroadcasterDropsSlowSubscriber(t *testing.T) {
	broadcaster := NewBroadcaster[int](1)
	fast, cancelFast := broadcaster.Subscribe()
	defer cancelFast()
	_, cancelSlow := broadcaster.Subscribe()
	defer cancelSlow()
	broadcaster.Publish(1)
	if got := <-fast; got != 1 {
		t.Fatalf("first message = %d, want 1", got)
	}
	broadcaster.Publish(2)
	if got := <-fast; got != 2 {
		t.Fatalf("second message = %d, want 2", got)
	}
}
