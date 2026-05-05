package main

import (
	"log"

	"zeaz.dev/meta/platform/internal/app"
	"zeaz.dev/meta/platform/internal/config"
	"zeaz.dev/meta/platform/internal/server"
)

func main() {
	const serviceName = "event-bus"
	cfg := config.Load(serviceName)
	routes := app.Routes(serviceName, []app.EndpointSpec{
		{Method: "POST", Path: "/v1/events", Scope: "event:publish", Action: "event.publish"},
		{Method: "GET", Path: "/v1/events/subscriptions", Scope: "event:read", Action: "event.subscriptions.list"},
	})
	if err := server.Run(cfg, routes); err != nil {
		log.Fatal(err)
	}
}
