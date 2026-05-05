package main

import (
	"encoding/json"
	"log"
	"net/http"

	"zeaz.dev/meta/platform/internal/config"
	zcrypto "zeaz.dev/meta/platform/internal/crypto"
	"zeaz.dev/meta/platform/internal/server"
)

func main() {
	cfg := config.Load("tss-signer")
	signer := zcrypto.PolicyGatedThresholdSigner{}
	routes := []server.Route{{Method: "POST", Path: "/v1/tss/sign", Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req zcrypto.SigningRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp, err := signer.Sign(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})}}
	if err := server.Run(cfg, routes); err != nil {
		log.Fatal(err)
	}
}
