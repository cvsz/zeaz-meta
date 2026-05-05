package main

import (
	"log"

	"zeaz.dev/meta/platform/internal/app"
	"zeaz.dev/meta/platform/internal/config"
	"zeaz.dev/meta/platform/internal/server"
)

func main() {
	const serviceName = "api-gateway"
	cfg := config.Load(serviceName)
	routes := app.Routes(serviceName, []app.EndpointSpec{
		{Method: "GET", Path: "/v1/gateway/routes", Scope: "gateway:read", Action: "gateway.routes.list"},
		{Method: "POST", Path: "/v1/gateway/canary", Scope: "gateway:deploy", Action: "gateway.canary.promote"},
	})
	if err := server.Run(cfg, routes); err != nil {
		log.Fatal(err)
	}
}
