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

	for currentlevel <= 2 {
		var previousMissionHistory, previousScenarioHistory, currentEnemyName = AssignMissions.GeneratedMission(previousMission, previousScenario)
		var enemyAttack = int(userCharacterSelected.defense - (userCharacterSelected.defense * 90 / 100))
		var enemyDefense = int(userCharacterSelected.attack - (userCharacterSelected.attack * 50 / 100))
		currentEnemy = NewCharacter(currentEnemyName, enemyAttack, enemyDefense, currentlevel)
		fmt.Println(currentEnemy, "\n")
		currentlevel++
		previousMission = previousMissionHistory
		previousScenario = previousScenarioHistory
	}

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
