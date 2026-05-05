package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"zeaz.dev/meta/platform/internal/config"
)

type Route struct {
	Method  string
	Path    string
	Handler http.Handler
}

func Run(cfg config.Config, routes []Route) error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"service": cfg.ServiceName, "status": "ok", "time": time.Now().UTC()})
	})
	mux.Handle("GET /metrics", promhttp.Handler())
	for _, route := range routes {
		mux.Handle(route.Method+" "+route.Path, route.Handler)
	}
	srv := &http.Server{Addr: cfg.HTTPAddr, Handler: tenantMiddleware(logMiddleware(mux)), ReadHeaderTimeout: cfg.ReadTimeout}
	slog.Info("starting service", "service", cfg.ServiceName, "addr", cfg.HTTPAddr)
	return srv.ListenAndServe()
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		slog.Info("request", "method", r.Method, "path", r.URL.Path, "tenant", r.Header.Get("X-Tenant-ID"), "duration_ms", time.Since(start).Milliseconds())
	})
}

func tenantMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/healthz" && r.URL.Path != "/metrics" && r.Header.Get("X-Tenant-ID") == "" {
			http.Error(w, "missing tenant context", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
