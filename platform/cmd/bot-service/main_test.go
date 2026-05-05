package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"zeaz.dev/meta/platform/internal/integration"
)

func TestSignedWebhookFailsClosedWithoutProviderSecret(t *testing.T) {
	t.Setenv("LINE_CHANNEL_SECRET", "")
	replayCache = integration.NewReplayCache(time.Minute, 10)
	req := httptest.NewRequest(http.MethodPost, "/v1/bots/line/webhook", strings.NewReader(`{"events":[]}`))
	req.Header.Set("X-Tenant-ID", "tenant-a")
	req.Header.Set("X-User-ID", "user-a")
	req.Header.Add("X-Zeaz-Scope", "bot:webhook")
	req.Header.Set("X-Line-Signature", "attacker-empty-secret-signature")
	rr := httptest.NewRecorder()
	signedWebhook("line").ServeHTTP(rr, req)
	if rr.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected fail-closed service unavailable, got %d", rr.Code)
	}
}

func TestSignedWebhookAcceptsValidLineSignature(t *testing.T) {
	secret := "line-secret"
	t.Setenv("LINE_CHANNEL_SECRET", secret)
	replayCache = integration.NewReplayCache(time.Minute, 10)
	body := []byte(`{"events":[]}`)
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write(body)
	req := httptest.NewRequest(http.MethodPost, "/v1/bots/line/webhook", strings.NewReader(string(body)))
	req.Header.Set("X-Tenant-ID", "tenant-a")
	req.Header.Set("X-User-ID", "user-a")
	req.Header.Add("X-Zeaz-Scope", "bot:webhook")
	req.Header.Set("X-Line-Signature", base64.StdEncoding.EncodeToString(mac.Sum(nil)))
	rr := httptest.NewRecorder()
	signedWebhook("line").ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected valid webhook, got %d: %s", rr.Code, rr.Body.String())
	}
}
