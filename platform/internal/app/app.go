package app

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"zeaz.dev/meta/platform/internal/audit"
	"zeaz.dev/meta/platform/internal/rbac"
	"zeaz.dev/meta/platform/internal/server"
)

type EndpointSpec struct {
	Method string
	Path   string
	Scope  string
	Action string
}

type Response struct {
	OK           bool      `json:"ok"`
	Service      string    `json:"service"`
	Action       string    `json:"action"`
	AuditEventID string    `json:"audit_event_id"`
	Time         time.Time `json:"time"`
}

func Routes(service string, specs []EndpointSpec) []server.Route {
	routes := make([]server.Route, 0, len(specs))
	for _, spec := range specs {
		s := spec
		routes = append(routes, server.Route{Method: s.Method, Path: s.Path, Handler: Handler(service, s.Scope, s.Action)})
	}
	return routes
}

func Handler(service, requiredScope, action string) http.Handler {
	respond := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		evt := audit.NewEvent(r.Header.Get("X-Tenant-ID"), r.Header.Get("X-User-ID"), action, r.URL.Path, map[string]any{"method": r.Method})
		WriteJSON(w, http.StatusOK, Response{OK: true, Service: service, Action: action, AuditEventID: evt.ID, Time: time.Now().UTC()})
	})
	if strings.HasPrefix(requiredScope, "auth:public") {
		return respond
	}
	return rbac.Require(requiredScope, respond)
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
