package AssignMissions

import (
	"fmt"
	"math/rand"
)

// generateMission genera una misión al azar y la devuelve como string.
func GeneratedMission(previousMission string, previousScenario string) (string, string, string) {
	// Lista de posibles tipos de misiones, objetivos, escenarios y enemigos.
	missionTypes := []string{
		"Exploración de un territorio desconocido",
		"Rescate de un equipo de exploradores perdidos",
		"Eliminación de una amenaza desconocida",
		"Recolección de recursos vitales para la supervivencia",
	}
	objectives := []string{
		"recuperando un artefacto valioso",
		"localizando un refugio seguro",
		"neutralizando una criatura peligrosa",
		"encontrando suministros médicos esenciales",
	}
	scenarios := []string{
		"en un bosque oscuro",
		"en una montaña nevada",
		"en un pantano peligroso",
		"en un desierto implacable",
	}
	enemies := []string{
		"lobos feroces",
		"orcos sanguinarios",
		"aracnidos venenosos",
		"espectros sombríos",
	}

	// Selecciona aleatoriamente un tipo de misión, un objetivo, un escenario y un enemigo.
	missionType := missionTypes[rand.Intn(len(missionTypes))]
	objective := objectives[rand.Intn(len(objectives))]
	scenario := scenarios[rand.Intn(len(scenarios))]
	enemy := enemies[rand.Intn(len(enemies))]

	// Construye la descripción de la misión.
	mission := "¡Atención! Misión de " + missionType + " " + objective + " " + scenario + " " + ". Enfrenta y derrota a los" + " " + enemy

	// Agrega una conexión con la misión anterior si es posible.
	if previousMission != "" {
		mission = "Después de completar la misión anterior, ahora te enfrentas a la siguiente tarea:" + mission
	}

	fmt.Println(mission)

	return mission, scenario, enemy
}
