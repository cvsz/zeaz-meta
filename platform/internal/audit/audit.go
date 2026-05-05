package audit

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

type Event struct {
	ID        string         `json:"id"`
	TenantID  string         `json:"tenant_id"`
	ActorID   string         `json:"actor_id"`
	Action    string         `json:"action"`
	Resource  string         `json:"resource"`
	Metadata  map[string]any `json:"metadata"`
	Hash      string         `json:"hash"`
	CreatedAt time.Time      `json:"created_at"`
}

func NewEvent(tenantID, actorID, action, resource string, metadata map[string]any) Event {
	now := time.Now().UTC()
	raw, _ := json.Marshal([]any{tenantID, actorID, action, resource, metadata, now.Format(time.RFC3339Nano)})
	sum := sha256.Sum256(raw)
	id := hex.EncodeToString(sum[:16])
	return Event{ID: id, TenantID: tenantID, ActorID: actorID, Action: action, Resource: resource, Metadata: metadata, Hash: hex.EncodeToString(sum[:]), CreatedAt: now}
}
