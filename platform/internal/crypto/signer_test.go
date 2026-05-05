package crypto

import "testing"

func TestPolicyGatedThresholdSignerRequiresPolicyFields(t *testing.T) {
	_, err := (PolicyGatedThresholdSigner{}).Sign(SigningRequest{TenantID: "tenant-a", KeyID: "key-a", Payload: []byte("payload")}, ThresholdSignatureProof{})
	if err == nil {
		t.Fatal("expected missing policy and nonce to fail")
	}
}

func TestPolicyGatedThresholdSignerFailsClosedWithoutRealQuorumProof(t *testing.T) {
	_, err := (PolicyGatedThresholdSigner{}).Sign(SigningRequest{TenantID: "tenant-a", KeyID: "key-a", PolicyID: "policy-a", Nonce: "nonce-a", Payload: []byte("payload")}, ThresholdSignatureProof{})
	if err == nil {
		t.Fatal("expected missing threshold proof to fail")
	}
}

func TestPolicyGatedThresholdSignerAcceptsQuorumTranscript(t *testing.T) {
	req := SigningRequest{TenantID: "tenant-a", KeyID: "key-a", PolicyID: "policy-a", Nonce: "nonce-a", Payload: []byte("payload")}
	proof := ThresholdSignatureProof{Scheme: "frost-ed25519", Threshold: 2, AggregateSignature: "sig:aggregate", GroupPublicKey: "pub:group", ProtocolTranscriptID: "session-1"}
	hash := TranscriptHash(req, proof)
	proof.Participants = []ParticipantTranscript{
		{ParticipantID: "p1", SpiffeID: "spiffe://zeaz.dev/ns/zeaz-signer/sa/signer-a", Round: 2, TranscriptHash: hash},
		{ParticipantID: "p2", SpiffeID: "spiffe://zeaz.dev/ns/zeaz-signer/sa/signer-b", Round: 2, TranscriptHash: hash},
	}
	resp, err := (PolicyGatedThresholdSigner{}).Sign(req, proof)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Digest == "" || resp.Signature != proof.AggregateSignature || resp.Scheme != "frost-ed25519" || resp.ProtocolTranscriptID != "session-1" {
		t.Fatalf("unexpected response: %#v", resp)
	}
}
