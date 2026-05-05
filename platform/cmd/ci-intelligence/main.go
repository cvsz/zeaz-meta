package main

import (
	"log"

	"zeaz.dev/meta/platform/internal/app"
	"zeaz.dev/meta/platform/internal/config"
	"zeaz.dev/meta/platform/internal/server"
)

func main() {
	const serviceName = "ci-intelligence"
	cfg := config.Load(serviceName)
	routes := app.Routes(serviceName, []app.EndpointSpec{
		{Method: "POST", Path: "/v1/ci/analyze", Scope: "ci:analyze", Action: "ci.analyze"},
		{Method: "POST", Path: "/v1/ci/rollback", Scope: "ci:rollback", Action: "ci.rollback"},
	})
	if err := server.Run(cfg, routes); err != nil {
		log.Fatal(err)
	}
}
