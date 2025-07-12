// Package hub provides a client for interacting with the hub service.
package hub

import (
	"testing"
	"time"
)

func TestNewHub(t *testing.T) {
	hub := newHub()
	if hub == nil {
		t.Fatal("expected new Hub instance, got nil")
	}
	if len(hub.clients) != 0 {
		t.Errorf("expected empty clients map, got %d clients", len(hub.clients))
	}
	if hub.broadcast == nil || hub.register == nil || hub.unregister == nil {
		t.Error("expected non-nil channels for broadcast, register, and unregister")
	}
}

func TestHubRun(t *testing.T) {
	hub := newHub()
	go hub.run()

	client1 := &Client{hub: hub, send: make(chan []byte, 1)}
	client2 := &Client{hub: hub, send: make(chan []byte, 1)}

	hub.register <- client1
	hub.register <- client2
	time.Sleep(10 * time.Millisecond)

	if len(hub.clients) != 2 {
		t.Errorf("expected 2 clients registered, got %d", len(hub.clients))
	}

	hub.unregister <- client1
	time.Sleep(10 * time.Millisecond)
	if len(hub.clients) != 1 {
		t.Errorf("expected 1 client after unregistering, got %d", len(hub.clients))
	}

	hub.broadcast <- []byte("test message")
	msg := <-client2.send
	if string(msg) != "test message" {
		t.Errorf("expected 'test message', got '%s'", msg)
	}
}
