package main

import "fmt"

type Character struct {
	name    string
	attack  int
	defense int
	level   int
	role    string
}

func NewCharacter(name string, attack int, defense int, level int) Character {
	return Character{name: name, attack: attack, defense: defense, level: level, role: "enemy"}
}

func main() {
	Play()
}

func Play() {
	characterOptions := map[string]Character{
		"Legolas": NewCharacter("Legolas", 20, 10, 1),
		"Gandalf": NewCharacter("Gandalf", 30, 15, 1),
		"Aragorn": NewCharacter("Aragorn", 25, 20, 1),
		"Gollum":  NewCharacter("Gollum", 18, 8, 1),
	}

	var playerCharacter = "Legolas"
	if character, ok := (*&characterOptions)[playerCharacter]; ok {
		character.SetRole("player")
		(*&characterOptions)[playerCharacter] = character
	}

	fmt.Println(characterOptions)

}
func (CharacterToSet *Character) SetRole(role string) {
	CharacterToSet.role = role

}
func AddCharacter(characterOptions *map[string]Character, characterTocreate *Character) {
	(*characterOptions)[characterTocreate.name] = *characterTocreate
}
