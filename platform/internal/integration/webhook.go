package integration

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	ErrMissingSignature = errors.New("missing webhook signature")
	ErrInvalidSignature = errors.New("invalid webhook signature")
	ErrReplay           = errors.New("webhook replay detected")
)

type ReplayCache struct {
	mu       sync.Mutex
	ttl      time.Duration
	entries  map[string]time.Time
	now      func() time.Time
	maxItems int
}

func NewReplayCache(ttl time.Duration, maxItems int) *ReplayCache {
	if ttl <= 0 {
		ttl = 10 * time.Minute
	}
	if maxItems <= 0 {
		maxItems = 10000
	}
	return &ReplayCache{ttl: ttl, entries: map[string]time.Time{}, now: time.Now, maxItems: maxItems}
}

func (c *ReplayCache) SeenOrStore(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := c.now().UTC()
	for k, expires := range c.entries {
		if !expires.After(now) {
			delete(c.entries, k)
		}
	}
	if _, ok := c.entries[key]; ok {
		return true
	}
	if len(c.entries) >= c.maxItems {
		oldestKey := ""
		oldestExpiry := now.Add(c.ttl)
		for k, expires := range c.entries {
			if expires.Before(oldestExpiry) {
				oldestKey, oldestExpiry = k, expires
			}
		}
		if oldestKey != "" {
			delete(c.entries, oldestKey)
		}
	}
	c.entries[key] = now.Add(c.ttl)
	return false
}

func VerifyLINE(secret string, body []byte, signature string, cache *ReplayCache) error {
	if strings.TrimSpace(signature) == "" {
		return ErrMissingSignature
	}
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write(body)
	expected := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	if !hmac.Equal([]byte(expected), []byte(signature)) {
		return ErrInvalidSignature
	}
	if cache != nil && cache.SeenOrStore("line:"+expected) {
		return ErrReplay
	}
	return nil
}

func VerifyTikTok(secret string, body []byte, signature string, headers http.Header, cache *ReplayCache) error {
	if strings.TrimSpace(signature) == "" {
		return ErrMissingSignature
	}
	canonicalHeaders := make([]string, 0, len(headers))
	for key, values := range headers {
		lower := strings.ToLower(key)
		if strings.HasPrefix(lower, "x-tt-") || lower == "content-type" {
			canonicalHeaders = append(canonicalHeaders, lower+":"+strings.Join(values, ","))
		}
	}
	sort.Strings(canonicalHeaders)
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write([]byte(strings.Join(canonicalHeaders, "\n")))
	_, _ = mac.Write([]byte("\n"))
	_, _ = mac.Write(body)
	expected := hex.EncodeToString(mac.Sum(nil))
	provided := strings.TrimPrefix(strings.ToLower(strings.TrimSpace(signature)), "sha256=")
	if !hmac.Equal([]byte(expected), []byte(provided)) {
		return ErrInvalidSignature
	}
	if cache != nil && cache.SeenOrStore("tiktok:"+provided) {
		return ErrReplay
	}
	return nil
}
