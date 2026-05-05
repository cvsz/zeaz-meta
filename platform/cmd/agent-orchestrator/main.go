package main

import (
	"encoding/json"
	"log"
	"net/http"

	"zeaz.dev/meta/platform/internal/agents"
	"zeaz.dev/meta/platform/internal/app"
	"zeaz.dev/meta/platform/internal/config"
	"zeaz.dev/meta/platform/internal/rbac"
	"zeaz.dev/meta/platform/internal/server"
)

type planRequest struct {
	Agents []agents.Agent `json:"agents"`
	Tasks  []agents.Task  `json:"tasks"`
}

func main() {
	const serviceName = "agent-orchestrator"
	cfg := config.Load(serviceName)
	routes := app.Routes(serviceName, []app.EndpointSpec{
		{Method: "POST", Path: "/v1/agents", Scope: "agent:create", Action: "agent.create"},
		{Method: "GET", Path: "/v1/agents", Scope: "agent:read", Action: "agent.list"},
	})
	routes = append(routes, server.Route{Method: "POST", Path: "/v1/agents/plans", Handler: rbac.Require("agent:plan", http.HandlerFunc(plan))})
	if err := server.Run(cfg, routes); err != nil {
		log.Fatal(err)
	}
}

func plan(w http.ResponseWriter, r *http.Request) {
	var req planRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	orch := agents.NewOrchestrator()
	for _, agent := range req.Agents {
		if err := orch.Register(agent); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
	}
	ordered, err := orch.Plan(req.Tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	app.WriteJSON(w, http.StatusOK, map[string]any{"ok": true, "ordered_tasks": ordered})
}
