package ledger

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

type Entry struct {
	AccountID string `json:"account_id"`
	Currency  string `json:"currency"`
	Amount    int64  `json:"amount"`
}

type Journal struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenant_id"`
	ReferenceID string    `json:"reference_id"`
	Entries     []Entry   `json:"entries"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewJournal(tenantID, referenceID string, entries []Entry) (Journal, error) {
	if tenantID == "" || referenceID == "" {
		return Journal{}, errors.New("tenant_id and reference_id are required")
	}
	if len(entries) < 2 {
		return Journal{}, errors.New("double-entry journal requires at least two entries")
	}
	balances := map[string]int64{}
	for _, entry := range entries {
		if entry.AccountID == "" || entry.Currency == "" || entry.Amount == 0 {
			return Journal{}, errors.New("each entry requires account_id, currency, and non-zero amount")
		}
		balances[entry.Currency] += entry.Amount
	}
	for currency, balance := range balances {
		if balance != 0 {
			return Journal{}, fmt.Errorf("journal is not balanced for %s", currency)
		}
	}
	now := time.Now().UTC()
	h := sha256.Sum256([]byte(fmt.Sprintf("%s|%s|%v|%s", tenantID, referenceID, entries, now.Format(time.RFC3339Nano))))
	return Journal{ID: hex.EncodeToString(h[:16]), TenantID: tenantID, ReferenceID: referenceID, Entries: entries, CreatedAt: now}, nil
}
