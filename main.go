package main

import (
	"escape-room-challenge/maps"
	"escape-room-challenge/system"
	"escape-room-challenge/unit"
	"escape-room-challenge/utils"
)

func main() {

	char := unit.NewCharacter()
	char.SetName()

	stage := maps.GetStage2()

	systemMessage := ""

	for {
		if stage.GetThisLocation().IsGoal {
			utils.PrintWIN()
			break
		}

		system.Print(&systemMessage, stage, char)
		system.Act(&char, &systemMessage, &stage)
	}

}
