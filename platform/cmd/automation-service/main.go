package main

import (
	"log"

	"zeaz.dev/meta/platform/internal/app"
	"zeaz.dev/meta/platform/internal/config"
	"zeaz.dev/meta/platform/internal/server"
)

func main() {
	const serviceName = "automation-service"
	cfg := config.Load(serviceName)
	routes := app.Routes(serviceName, []app.EndpointSpec{
		{Method: "POST", Path: "/v1/automation/line/jobs", Scope: "automation:write", Action: "automation.line.job.create"},
		{Method: "POST", Path: "/v1/automation/tiktok/jobs", Scope: "automation:write", Action: "automation.tiktok.job.create"},
	})
	if err := server.Run(cfg, routes); err != nil {
		log.Fatal(err)
	}
}
