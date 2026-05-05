package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"zeaz.dev/meta/platform/internal/audit"
	"zeaz.dev/meta/platform/internal/config"
	"zeaz.dev/meta/platform/internal/integration"
	"zeaz.dev/meta/platform/internal/rbac"
	"zeaz.dev/meta/platform/internal/server"
)

var replayCache = integration.NewReplayCache(10*time.Minute, 50000)

func main() {
	cfg := config.Load("bot-service")
	routes := []server.Route{
		{Method: "POST", Path: "/v1/bots/line/webhook", Handler: signedWebhook("line")},
		{Method: "POST", Path: "/v1/bots/tiktok/webhook", Handler: signedWebhook("tiktok")},
	}
	if err := server.Run(cfg, routes); err != nil {
		log.Fatal(err)
	}
}

func endpoint(requiredScope, action string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(requiredScope, "auth:public") {
			rbac.Require(requiredScope, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { respond(w, r, action, nil) })).ServeHTTP(w, r)
			return
		}
		respond(w, r, action, nil)
	})
}

func signedWebhook(provider string) http.Handler {
	return rbac.Require("bot:webhook", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(http.MaxBytesReader(w, r.Body, 1<<20))
		if err != nil {
			http.Error(w, "webhook body too large", http.StatusRequestEntityTooLarge)
			return
		}
		var verifyErr error
		switch provider {
		case "line":
			secret := os.Getenv("LINE_CHANNEL_SECRET")
			if secret == "" {
				http.Error(w, "LINE webhook secret is unavailable", http.StatusServiceUnavailable)
				return
			}
			verifyErr = integration.VerifyLINE(secret, body, r.Header.Get("X-Line-Signature"), replayCache)
		case "tiktok":
			secret := os.Getenv("TIKTOK_WEBHOOK_SECRET")
			if secret == "" {
				http.Error(w, "TikTok webhook secret is unavailable", http.StatusServiceUnavailable)
				return
			}
			verifyErr = integration.VerifyTikTok(secret, body, r.Header.Get("X-Tt-Signature"), r.Header, replayCache)
		default:
			verifyErr = errors.New("unsupported provider")
		}
		if verifyErr != nil {
			http.Error(w, verifyErr.Error(), http.StatusUnauthorized)
			return
		}
		respond(w, r, "bot."+provider+".webhook", map[string]any{"provider": provider, "payload_bytes": len(body)})
	}))
}

func respond(w http.ResponseWriter, r *http.Request, action string, metadata map[string]any) {
	if metadata == nil {
		metadata = map[string]any{}
	}
	metadata["method"] = r.Method
	evt := audit.NewEvent(r.Header.Get("X-Tenant-ID"), r.Header.Get("X-User-ID"), action, r.URL.Path, metadata)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"ok": true, "service": "bot-service", "action": action, "audit_event_id": evt.ID, "time": time.Now().UTC()})
}
