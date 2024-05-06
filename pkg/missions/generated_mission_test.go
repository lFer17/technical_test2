package AssignMissions_test

import (
	AssignMissions "lfer17/user/test/pkg/missions"
	"testing"
)

func TestGeneratedMission(t *testing.T) {
	// test when previousMission & previousScenario are empty
	mission, scenario, enemy := AssignMissions.GeneratedMission("", "")

	if mission == "" || scenario == "" || enemy == "" {
		t.Error("La misión, el escenario o el enemigo están vacíos cuando no deberían estarlo")
	}
	// test when previousMission & previousScenario has content
	previousMission := "Misión anterior"
	previousScenario := "Escenario anterior"
	mission, scenario, enemy = AssignMissions.GeneratedMission(previousMission, previousScenario)

	if mission == "" || scenario == "" || enemy == "" {
		t.Error("La misión, el escenario o el enemigo están vacíos cuando no deberían estarlo")
	}
}
