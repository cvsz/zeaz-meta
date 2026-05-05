package rbac

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequireAllowsTenantUserAndScope(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/audit/events", nil)
	req.Header.Set("X-Tenant-ID", "tenant-a")
	req.Header.Set("X-User-ID", "user-a")
	req.Header.Add("X-Zeaz-Scope", "audit:read")
	rr := httptest.NewRecorder()
	Require("audit:read", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNoContent) })).ServeHTTP(rr, req)
	if rr.Code != http.StatusNoContent {
		t.Fatalf("expected allowed request, got %d", rr.Code)
	}
}

func TestRequireDeniesMissingScope(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/audit/events", nil)
	req.Header.Set("X-Tenant-ID", "tenant-a")
	req.Header.Set("X-User-ID", "user-a")
	rr := httptest.NewRecorder()
	Require("audit:read", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { t.Fatal("handler should not run") })).ServeHTTP(rr, req)
	if rr.Code != http.StatusForbidden {
		t.Fatalf("expected forbidden, got %d", rr.Code)
	}
}
