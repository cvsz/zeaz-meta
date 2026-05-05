package events

import (
	"encoding/json"
	"testing"
)

func TestNewEventValidatesPayload(t *testing.T) {
	event, err := New("tenant-a", "wallet.created", json.RawMessage(`{"id":"wallet-1"}`))
	if err != nil {
		t.Fatal(err)
	}
	if event.ID == "" || event.Subject != "wallet.created" {
		t.Fatalf("unexpected event: %#v", event)
	}
	if _, err := New("tenant-a", "wallet.created", json.RawMessage(`{`)); err == nil {
		t.Fatal("expected invalid json rejection")
	}
}
