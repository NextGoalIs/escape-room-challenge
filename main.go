package main

import (
	"escape-room-challenge/system"
	"escape-room-challenge/unit"
)

func main() {

	char := unit.NewCharacter()

	system.PlayStage1(char)
	system.PlayStage2(char)

}
