package main

import (
	"log"

	"zeaz.dev/meta/platform/internal/app"
	"zeaz.dev/meta/platform/internal/config"
	"zeaz.dev/meta/platform/internal/server"
)

func main() {
	const serviceName = "scheduler-service"
	cfg := config.Load(serviceName)
	routes := app.Routes(serviceName, []app.EndpointSpec{
		{Method: "POST", Path: "/v1/schedules", Scope: "schedule:write", Action: "scheduler.schedule.create"},
		{Method: "GET", Path: "/v1/schedules", Scope: "schedule:read", Action: "scheduler.schedules.list"},
	})
	if err := server.Run(cfg, routes); err != nil {
		log.Fatal(err)
	}
}
