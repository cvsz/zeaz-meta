package main

import (
	"log"

	"zeaz.dev/meta/platform/internal/app"
	"zeaz.dev/meta/platform/internal/config"
	"zeaz.dev/meta/platform/internal/server"
)

func main() {
	const serviceName = "refactor-engine"
	cfg := config.Load(serviceName)
	routes := app.Routes(serviceName, []app.EndpointSpec{
		{Method: "POST", Path: "/v1/refactor/plans", Scope: "refactor:plan", Action: "refactor.plan.create"},
		{Method: "POST", Path: "/v1/refactor/apply", Scope: "refactor:apply", Action: "refactor.apply"},
	})
	if err := server.Run(cfg, routes); err != nil {
		log.Fatal(err)
	}
}
