package crypto

import "testing"

func TestPolicyGatedThresholdSignerRequiresPolicyFields(t *testing.T) {
	_, err := (PolicyGatedThresholdSigner{}).Sign(SigningRequest{TenantID: "tenant-a", KeyID: "key-a", Payload: []byte("payload")})
	if err == nil {
		t.Fatal("expected missing policy and nonce to fail")
	}
}

func TestPolicyGatedThresholdSignerReturnsDigestEnvelope(t *testing.T) {
	resp, err := (PolicyGatedThresholdSigner{}).Sign(SigningRequest{TenantID: "tenant-a", KeyID: "key-a", PolicyID: "policy-a", Nonce: "nonce-a", Payload: []byte("payload")})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Digest == "" || resp.Signature == "" || resp.Scheme != "policy-gated-threshold-interface" {
		t.Fatalf("unexpected response: %#v", resp)
	}
}
