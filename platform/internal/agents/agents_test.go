package agents

import "testing"

func TestOrchestratorPlansDAG(t *testing.T) {
	orch := NewOrchestrator()
	for _, agent := range []Agent{{ID: "planner-1", Role: Planner, Capabilities: []string{"plan"}}, {ID: "executor-1", Role: Executor, Capabilities: []string{"apply"}}} {
		if err := orch.Register(agent); err != nil {
			t.Fatal(err)
		}
	}
	ordered, err := orch.Plan([]Task{{ID: "apply", AgentID: "executor-1", DependsOn: []string{"plan"}}, {ID: "plan", AgentID: "planner-1"}})
	if err != nil {
		t.Fatal(err)
	}
	if ordered[0].ID != "plan" || ordered[1].ID != "apply" {
		t.Fatalf("unexpected order: %#v", ordered)
	}
}

func TestOrchestratorRejectsCycles(t *testing.T) {
	orch := NewOrchestrator()
	if err := orch.Register(Agent{ID: "planner-1", Role: Planner, Capabilities: []string{"plan"}}); err != nil {
		t.Fatal(err)
	}
	_, err := orch.Plan([]Task{{ID: "a", AgentID: "planner-1", DependsOn: []string{"b"}}, {ID: "b", AgentID: "planner-1", DependsOn: []string{"a"}}})
	if err == nil {
		t.Fatal("expected cycle rejection")
	}
}
