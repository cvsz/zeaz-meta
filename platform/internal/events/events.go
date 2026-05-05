package events

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"
)

type Event struct {
	ID        string          `json:"id"`
	TenantID  string          `json:"tenant_id"`
	Subject   string          `json:"subject"`
	Payload   json.RawMessage `json:"payload"`
	CreatedAt time.Time       `json:"created_at"`
}

func New(tenantID, subject string, payload json.RawMessage) (Event, error) {
	if tenantID == "" || subject == "" || !json.Valid(payload) {
		return Event{}, errors.New("tenant_id, subject, and valid json payload are required")
	}
	now := time.Now().UTC()
	h := sha256.Sum256(append(append([]byte(tenantID+"|"+subject+"|"), payload...), []byte(now.Format(time.RFC3339Nano))...))
	return Event{ID: hex.EncodeToString(h[:16]), TenantID: tenantID, Subject: subject, Payload: payload, CreatedAt: now}, nil
}
