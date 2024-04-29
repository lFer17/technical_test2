package main

import (
	"fmt"
	AssignMissions "lfer17/user/test/pkg/missions"
)

type Character struct {
	name    string
	attack  int
	defense int
	level   int
	role    string
}

var (
	userCharacterSelected Character
	currentEnemy          Character
	currentlevel          int
)

func NewCharacter(name string, attack int, defense int, level int) Character {
	return Character{name: name, attack: attack, defense: defense, level: level, role: "enemy"}
}

func main() {

	Play()
}

func Play() {
	SettingUserCharacter()
	var previousMission string
	var previousScenario string

	// loop de jugabilidad inifinita
	// Generating misions and history
	var _, _, currentEnemyName = AssignMissions.GeneratedMission(previousMission, previousScenario)

	// Creating enemy capacity
	var enemyAttack = int(userCharacterSelected.defense - (userCharacterSelected.defense * 90 / 100))
	var enemyDefense = int(userCharacterSelected.attack + (userCharacterSelected.attack * 20 / 100))
	currentEnemy = NewCharacter(currentEnemyName, enemyAttack, enemyDefense, currentlevel)

	// increasing  level
	currentlevel++

	fmt.Printf(" ha aparecido %s \n", currentEnemy.name)

	fmt.Println("entramos en combate? ingresa : si/no")
	var userResponse string
	fmt.Scan(&userResponse)

	if userResponse == "si" {
		CombatInit(&userCharacterSelected, &currentEnemy)
	} else {
		EscapeOptions()
	}

	// assigning history and mission to mantain narrative consistency
	// previousMission = previousMissionHistory
	// previousScenario = previousScenarioHistory

}

// Despliega opciones para que el usuario escoja su personaje ingresando un entero
func SettingUserCharacter() {
	characterOptions := map[string]Character{
		"1": NewCharacter("Legolas", 20, 10, 1),
		"2": NewCharacter("Gandalf", 30, 15, 1),
		"3": NewCharacter("Aragorn", 25, 20, 1),
		"4": NewCharacter("Gollum", 18, 8, 1),
	}
	fmt.Println(" Ingresa un numero para elegir tu personaje \n")

	for name, character := range characterOptions {
		fmt.Println(name)
		fmt.Printf("Nombre: %s\n", character.name)
		fmt.Printf("Ataque: %d\n", character.attack)
		fmt.Printf("Defensa: %d\n\n", character.defense)
	}

	var userCharacterSelection string

	fmt.Scanln(&userCharacterSelection)

	switch userCharacterSelection {
	case "1":
		userCharacterSelected = characterOptions["1"]
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Vamos allá ", userCharacterSelected.name)
		fmt.Println("-----------------------------------------------------")
	case "2":
		userCharacterSelected = characterOptions["2"]
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Vamos allá ", userCharacterSelected.name)
		fmt.Println("-----------------------------------------------------")
	case "3":
		userCharacterSelected = characterOptions["3"]
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Vamos allá ", userCharacterSelected.name)
		fmt.Println("-----------------------------------------------------")
	case "4":
		userCharacterSelected = characterOptions["4"]
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Vamos allá ", userCharacterSelected.name)
		fmt.Println("-----------------------------------------------------")
	}
}

// start combat between enemy and user
func CombatInit(currentCharacterParam *Character, currentEnemyParam *Character) {
	healthReference := currentCharacterParam.defense
	currentEnemyParam.defense = currentEnemyParam.defense - currentCharacterParam.attack
	currentCharacterParam.defense = currentCharacterParam.defense - currentEnemyParam.attack

	if currentEnemyParam.defense <= 0 {
		fmt.Println("Felicitaciones eliminaste al enemigo\n")
		fmt.Printf(" Se ha Restaurado tu salud %d \n", healthReference-1)
		// restoring user health
		currentCharacterParam.defense = healthReference - 1
	} else if currentEnemyParam.defense > 0 && currentCharacterParam.defense > 0 {
		fmt.Printf("haz dañado a tu enemigo con: %d \n", currentCharacterParam.attack)
		fmt.Printf(" defensa restante:%d \n", currentEnemyParam.defense)
		// dicreasing player defense
		fmt.Printf(" Enemigo te resta:%d \n", currentCharacterParam.defense)
		//  reiniciando combate
		fmt.Println(" volvemos atacar? : si / no")
		var userResponse string
		fmt.Scan(&userResponse)
		if userResponse == "si" && currentCharacterParam.defense > 0 {
			CombatInit(currentCharacterParam, currentEnemyParam)
		} else {
			EscapeOptions()
		}
	} else {
		fmt.Println("Aqui murio un valiente, Quedaste bien muertito!\n")
		fmt.Println(" Game over\n")
		GameOver()
	}

}

func EscapeOptions() {
	// generate option to run
	fmt.Println("bienvenid@, bitch pudiste escapar de: ", currentEnemy.name)

}

func GameOver() bool {
	return true
}
