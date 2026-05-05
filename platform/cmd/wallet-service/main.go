package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"zeaz.dev/meta/platform/internal/audit"
	"zeaz.dev/meta/platform/internal/config"
	"zeaz.dev/meta/platform/internal/rbac"
	"zeaz.dev/meta/platform/internal/server"
)

func main() {
	cfg := config.Load("wallet-service")
	routes := []server.Route{
		{Method: "GET", Path: "/v1/wallet/accounts", Handler: endpoint("wallet:read", "wallet.accounts.list")},
		{Method: "POST", Path: "/v1/wallet/transfers", Handler: endpoint("wallet:transfer:create", "wallet.transfer.create")},
	}
	if err := server.Run(cfg, routes); err != nil {
		log.Fatal(err)
	}
}

func endpoint(requiredScope, action string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(requiredScope, "auth:public") {
			rbac.Require(requiredScope, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { respond(w, r, action) })).ServeHTTP(w, r)
			return
		}
		respond(w, r, action)
	})
}

func respond(w http.ResponseWriter, r *http.Request, action string) {
	evt := audit.NewEvent(r.Header.Get("X-Tenant-ID"), r.Header.Get("X-User-ID"), action, r.URL.Path, map[string]any{"method": r.Method})
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"ok": true, "service": "wallet-service", "action": action, "audit_event_id": evt.ID, "time": time.Now().UTC()})
}
