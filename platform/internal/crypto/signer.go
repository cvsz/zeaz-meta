package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
)

// SigningRequest is the policy boundary accepted by tss-signer. Payload contains the canonical
// transaction bytes; callers never submit private keys or key shares.
type SigningRequest struct {
	TenantID string `json:"tenant_id"`
	KeyID    string `json:"key_id"`
	Payload  []byte `json:"payload"`
	PolicyID string `json:"policy_id"`
	Nonce    string `json:"nonce"`
}

// ParticipantTranscript is emitted by an isolated FROST/tss-lib participant after it validates the
// signing session, loads its share from Vault/HSM, completes the protocol rounds, and contributes to
// the aggregate signature. The coordinator verifies quorum metadata and transcript consistency.
type ParticipantTranscript struct {
	ParticipantID  string `json:"participant_id"`
	SpiffeID       string `json:"spiffe_id"`
	Round          uint8  `json:"round"`
	TranscriptHash string `json:"transcript_hash"`
}

// ThresholdSignatureProof is the aggregate TSS result returned by the isolated protocol workers.
type ThresholdSignatureProof struct {
	Scheme               string                  `json:"scheme"`
	Threshold            int                     `json:"threshold"`
	Participants         []ParticipantTranscript `json:"participants"`
	AggregateSignature   string                  `json:"aggregate_signature"`
	GroupPublicKey       string                  `json:"group_public_key"`
	ProtocolTranscriptID string                  `json:"protocol_transcript_id"`
}

type SigningResponse struct {
	Digest               string `json:"digest"`
	Signature            string `json:"signature"`
	Scheme               string `json:"scheme"`
	ProtocolTranscriptID string `json:"protocol_transcript_id"`
}

type Signer interface {
	Sign(req SigningRequest, proof ThresholdSignatureProof) (SigningResponse, error)
}

// PolicyGatedThresholdSigner is fail-closed coordination logic for real FROST/tss-lib workers. It
// deliberately does not synthesize signatures; production participants must provide a quorum proof.
type PolicyGatedThresholdSigner struct{}

func (PolicyGatedThresholdSigner) Sign(req SigningRequest, proof ThresholdSignatureProof) (SigningResponse, error) {
	if req.TenantID == "" || req.KeyID == "" || req.PolicyID == "" || req.Nonce == "" || len(req.Payload) == 0 {
		return SigningResponse{}, errors.New("tenant, key, policy, nonce, and payload are required")
	}
	if proof.Scheme != "frost-ed25519" && proof.Scheme != "tss-lib-ecdsa" {
		return SigningResponse{}, errors.New("unsupported threshold signing scheme")
	}
	if proof.Threshold < 2 || len(proof.Participants) < proof.Threshold {
		return SigningResponse{}, errors.New("threshold quorum was not satisfied")
	}
	if proof.AggregateSignature == "" || proof.GroupPublicKey == "" || proof.ProtocolTranscriptID == "" {
		return SigningResponse{}, errors.New("aggregate signature, group public key, and transcript id are required")
	}
	digest := sha256.Sum256(req.Payload)
	expectedTranscript := transcriptHash(req, hex.EncodeToString(digest[:]), proof.Scheme, proof.GroupPublicKey, proof.ProtocolTranscriptID)
	seen := map[string]bool{}
	for _, participant := range proof.Participants {
		if participant.ParticipantID == "" || participant.SpiffeID == "" || participant.Round == 0 || participant.TranscriptHash == "" {
			return SigningResponse{}, errors.New("participant transcript is incomplete")
		}
		if !strings.HasPrefix(participant.SpiffeID, "spiffe://zeaz.dev/ns/zeaz-signer/sa/") {
			return SigningResponse{}, errors.New("participant SPIFFE identity is outside signer trust domain")
		}
		if participant.TranscriptHash != expectedTranscript {
			return SigningResponse{}, errors.New("participant transcript hash mismatch")
		}
		if seen[participant.ParticipantID] {
			return SigningResponse{}, errors.New("duplicate participant transcript")
		}
		seen[participant.ParticipantID] = true
	}
	return SigningResponse{Digest: hex.EncodeToString(digest[:]), Signature: proof.AggregateSignature, Scheme: proof.Scheme, ProtocolTranscriptID: proof.ProtocolTranscriptID}, nil
}

// TranscriptHash computes the canonical hash that FROST/tss-lib participants must sign/attest.
func TranscriptHash(req SigningRequest, proof ThresholdSignatureProof) string {
	digest := sha256.Sum256(req.Payload)
	return transcriptHash(req, hex.EncodeToString(digest[:]), proof.Scheme, proof.GroupPublicKey, proof.ProtocolTranscriptID)
}

func transcriptHash(req SigningRequest, digest, scheme, groupPublicKey, transcriptID string) string {
	parts := []string{req.TenantID, req.KeyID, req.PolicyID, req.Nonce, digest, scheme, groupPublicKey, transcriptID}
	h := sha256.Sum256([]byte(strings.Join(parts, "|")))
	return hex.EncodeToString(h[:])
}
