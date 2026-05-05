package agents

import (
	"errors"
	"sort"
)

type AgentRole string

const (
	Planner      AgentRole = "planner"
	Executor     AgentRole = "executor"
	Reviewer     AgentRole = "reviewer"
	Architect    AgentRole = "architect"
	Economist    AgentRole = "economist"
	AgentFactory AgentRole = "agent_factory"
)

type Agent struct {
	ID           string    `json:"id"`
	Role         AgentRole `json:"role"`
	Capabilities []string  `json:"capabilities"`
	State        string    `json:"state"`
}

type Task struct {
	ID           string   `json:"id"`
	AgentID      string   `json:"agent_id"`
	DependsOn    []string `json:"depends_on"`
	RewardWeight int      `json:"reward_weight"`
}

type Orchestrator struct {
	agents map[string]Agent
}

func NewOrchestrator() *Orchestrator { return &Orchestrator{agents: map[string]Agent{}} }

func (o *Orchestrator) Register(agent Agent) error {
	if agent.ID == "" || agent.Role == "" || len(agent.Capabilities) == 0 {
		return errors.New("agent id, role, and capabilities are required")
	}
	if agent.State == "" {
		agent.State = "ready"
	}
	o.agents[agent.ID] = agent
	return nil
}

func (o *Orchestrator) Plan(tasks []Task) ([]Task, error) {
	ids := map[string]Task{}
	for _, task := range tasks {
		if task.ID == "" || task.AgentID == "" {
			return nil, errors.New("task id and agent id are required")
		}
		if _, ok := o.agents[task.AgentID]; !ok {
			return nil, errors.New("task references unknown agent")
		}
		ids[task.ID] = task
	}
	visiting, visited := map[string]bool{}, map[string]bool{}
	ordered := []Task{}
	var visit func(string) error
	visit = func(id string) error {
		if visiting[id] {
			return errors.New("task graph contains a cycle")
		}
		if visited[id] {
			return nil
		}
		task, ok := ids[id]
		if !ok {
			return errors.New("task dependency is missing")
		}
		visiting[id] = true
		deps := append([]string(nil), task.DependsOn...)
		sort.Strings(deps)
		for _, dep := range deps {
			if err := visit(dep); err != nil {
				return err
			}
		}
		visiting[id] = false
		visited[id] = true
		ordered = append(ordered, task)
		return nil
	}
	keys := make([]string, 0, len(ids))
	for id := range ids {
		keys = append(keys, id)
	}
	sort.Strings(keys)
	for _, id := range keys {
		if err := visit(id); err != nil {
			return nil, err
		}
	}
	return ordered, nil
}
