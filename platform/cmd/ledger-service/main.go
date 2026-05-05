package main

import (
	"log"

	"zeaz.dev/meta/platform/internal/app"
	"zeaz.dev/meta/platform/internal/config"
	"zeaz.dev/meta/platform/internal/server"
)

func main() {
	const serviceName = "ledger-service"
	cfg := config.Load(serviceName)
	routes := app.Routes(serviceName, []app.EndpointSpec{
		{Method: "POST", Path: "/v1/ledger/journals", Scope: "ledger:write", Action: "ledger.journal.create"},
		{Method: "GET", Path: "/v1/ledger/accounts", Scope: "ledger:read", Action: "ledger.accounts.list"},
	})
	if err := server.Run(cfg, routes); err != nil {
		log.Fatal(err)
	}
}
