package main

import (
	"log"

	"zeaz.dev/meta/platform/internal/app"
	"zeaz.dev/meta/platform/internal/config"
	"zeaz.dev/meta/platform/internal/server"
)

func main() {
	const serviceName = "signer-service"
	cfg := config.Load(serviceName)
	routes := app.Routes(serviceName, []app.EndpointSpec{
		{Method: "POST", Path: "/v1/signer/sessions", Scope: "signer:session:create", Action: "signer.session.create"},
		{Method: "GET", Path: "/v1/signer/keys", Scope: "signer:key:read", Action: "signer.keys.list"},
	})
	if err := server.Run(cfg, routes); err != nil {
		log.Fatal(err)
	}
}
