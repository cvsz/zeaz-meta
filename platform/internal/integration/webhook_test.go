package integration

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"testing"
	"time"
)

func TestVerifyLINEAcceptsValidSignatureAndRejectsReplay(t *testing.T) {
	body := []byte(`{"events":[{"type":"message"}]}`)
	mac := hmac.New(sha256.New, []byte("line-secret"))
	_, _ = mac.Write(body)
	sig := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	cache := NewReplayCache(time.Minute, 10)
	if err := VerifyLINE("line-secret", body, sig, cache); err != nil {
		t.Fatal(err)
	}
	if err := VerifyLINE("line-secret", body, sig, cache); err != ErrReplay {
		t.Fatalf("expected replay rejection, got %v", err)
	}
}

func TestVerifyLINERejectsInvalidSignature(t *testing.T) {
	if err := VerifyLINE("line-secret", []byte("{}"), "invalid", nil); err != ErrInvalidSignature {
		t.Fatalf("expected invalid signature, got %v", err)
	}
}

func TestVerifyTikTokUsesCanonicalHeadersAndBody(t *testing.T) {
	body := []byte(`{"type":"order_status"}`)
	headers := http.Header{"X-Tt-Timestamp": []string{"1710000000"}, "Content-Type": []string{"application/json"}}
	mac := hmac.New(sha256.New, []byte("tt-secret"))
	_, _ = mac.Write([]byte("content-type:application/json\nx-tt-timestamp:1710000000"))
	_, _ = mac.Write([]byte("\n"))
	_, _ = mac.Write(body)
	sig := hex.EncodeToString(mac.Sum(nil))
	if err := VerifyTikTok("tt-secret", body, "sha256="+sig, headers, NewReplayCache(time.Minute, 10)); err != nil {
		t.Fatal(err)
	}
}
