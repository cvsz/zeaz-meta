package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

type SigningRequest struct {
	TenantID string `json:"tenant_id"`
	KeyID    string `json:"key_id"`
	Payload  []byte `json:"payload"`
	PolicyID string `json:"policy_id"`
	Nonce    string `json:"nonce"`
}

type SigningResponse struct {
	Digest    string `json:"digest"`
	Signature string `json:"signature"`
	Scheme    string `json:"scheme"`
}

type Signer interface {
	Sign(req SigningRequest) (SigningResponse, error)
}

type PolicyGatedThresholdSigner struct{}

func (PolicyGatedThresholdSigner) Sign(req SigningRequest) (SigningResponse, error) {
	if req.TenantID == "" || req.KeyID == "" || req.PolicyID == "" || req.Nonce == "" || len(req.Payload) == 0 {
		return SigningResponse{}, errors.New("tenant, key, policy, nonce, and payload are required")
	}
	digest := sha256.Sum256(req.Payload)
	// This deterministic envelope is intentionally not a private-key signature. It is the stable interface
	// boundary where production deployments plug in tss-lib/FROST participants backed by Vault/HSM shares.
	proof := sha256.Sum256(append([]byte(req.TenantID+":"+req.KeyID+":"+req.PolicyID+":"+req.Nonce+":"), digest[:]...))
	return SigningResponse{Digest: hex.EncodeToString(digest[:]), Signature: hex.EncodeToString(proof[:]), Scheme: "policy-gated-threshold-interface"}, nil
}
