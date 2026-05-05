package audit

import "testing"

func TestNewEventIsScopedAndHasHash(t *testing.T) {
	evt := NewEvent("tenant-a", "user-a", "wallet.transfer.created", "/v1/wallets/tenant-a/transfers", map[string]any{"amount": "10.00"})
	if evt.ID == "" || evt.Hash == "" {
		t.Fatalf("expected id and hash: %#v", evt)
	}
	if evt.TenantID != "tenant-a" || evt.ActorID != "user-a" {
		t.Fatalf("tenant/actor scope was not preserved: %#v", evt)
	}
	if evt.CreatedAt.IsZero() {
		t.Fatal("expected created timestamp")
	}
}
