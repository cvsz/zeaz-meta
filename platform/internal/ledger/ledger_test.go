package ledger

import "testing"

func TestNewJournalRequiresBalancedEntries(t *testing.T) {
	journal, err := NewJournal("tenant-a", "ref-1", []Entry{{AccountID: "cash", Currency: "USD", Amount: 100}, {AccountID: "revenue", Currency: "USD", Amount: -100}})
	if err != nil {
		t.Fatalf("expected balanced journal: %v", err)
	}
	if journal.ID == "" || journal.TenantID != "tenant-a" {
		t.Fatalf("unexpected journal: %#v", journal)
	}
	if _, err := NewJournal("tenant-a", "ref-2", []Entry{{AccountID: "cash", Currency: "USD", Amount: 100}, {AccountID: "fees", Currency: "USD", Amount: -1}}); err == nil {
		t.Fatal("expected unbalanced journal to fail")
	}
}
