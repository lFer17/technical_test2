package main

import (
	"fmt"
	AssignMissions "lfer17/user/test/pkg/missions"
	Utils "lfer17/user/test/pkg/utils"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
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
	numberOfEnemies       int
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

	// Generating misions and history
	var _, _, currentEnemyName = AssignMissions.GeneratedMission(previousMission, previousScenario)

	// Creating enemy
	CreateNewEnemy(currentEnemyName)
	// increasing  level
	currentlevel++
	// assigning number of enemies in current level
	numberOfEnemies = currentlevel * 2

	// Game infinity loop

	for {
		if numberOfEnemies > 0 {
			fmt.Printf("Ha aparecido un nuevo %s", currentEnemyName)
			CreateNewEnemy(currentEnemyName)
			Utils.TimerDelay()
			Utils.ClearConsole()
		} else {
			Utils.ClearConsole()
			currentlevel++
			numberOfEnemies = currentlevel * 2
			fmt.Printf(" bienvenido al nivel %d \n", currentlevel)
			fmt.Printf(" Tu nueva mision será...")
			var _, _, currentEnemyName = AssignMissions.GeneratedMission(previousMission, previousScenario)
			CreateNewEnemy(currentEnemyName)
		}

		fmt.Printf(" ha aparecido %s \n", currentEnemy.name)

		fmt.Println("entramos en combate? ingresa : si/no")
		var userResponse string
		fmt.Scan(&userResponse)

		if userResponse == "si" {
			CombatInit(&userCharacterSelected, &currentEnemy)
		} else {
			EscapeOptions()
		}
	}

}

// display options to choose a character
func SettingUserCharacter() {
	characterOptions := map[string]Character{
		"1": NewCharacter("Legolas", 20, 10, 1),
		"2": NewCharacter("Gandalf", 30, 15, 1),
		"3": NewCharacter("Aragorn", 25, 20, 1),
		"4": NewCharacter("Gollum", 18, 8, 1),
	}
	fmt.Println(" Ingresa un numero para elegir tu personaje \n")
	table := tablewriter.NewWriter(os.Stdout)
	for key, character := range characterOptions {
		table.SetHeader([]string{"Opcion", "Nombre", "Ataque", "Defensa"})
		attackString := strconv.Itoa(character.attack)
		defenseString := strconv.Itoa(character.defense)
		table.Append([]string{key, character.name, attackString, defenseString})
	}
	table.Render()
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
	Utils.TimerDelay()
	Utils.ClearConsole()
}

func CreateNewEnemy(enemyname string) {
	var enemyAttack = int(userCharacterSelected.defense - (userCharacterSelected.defense * 90 / 100))
	var enemyDefense = int(userCharacterSelected.attack + (userCharacterSelected.attack * 20 / 100))
	currentEnemy = NewCharacter(enemyname, enemyAttack, enemyDefense, currentlevel)

}

// start combat between enemy and user
func CombatInit(currentCharacterParam *Character, currentEnemyParam *Character) {
	healthReference := currentCharacterParam.defense
	currentEnemyParam.defense = currentEnemyParam.defense - currentCharacterParam.attack
	currentCharacterParam.defense = currentCharacterParam.defense - currentEnemyParam.attack

	if currentEnemyParam.defense <= 0 {
		// Aquí puedes escribir tu solución al conflicto
		fmt.Println("Felicitaciones, has eliminado al enemigo.\n")
	}
	{
		fmt.Printf(" Se ha Restaurado tu salud %d \n", healthReference-1)
		numberOfEnemies--
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
	fmt.Println("ingresa el numero 1 para finalizar partida")
	var userResponse string
	fmt.Scan(&userResponse)
	if userResponse == "1" {
		GameOver()
	}
}

func GameOver() {
	os.Exit(0)
}
