package system

import (
	"escape-room-challenge/maps"
	"escape-room-challenge/unit"
	"escape-room-challenge/utils"
)

func PlayStage1(char unit.Character) {
	stage := maps.GetStage1()

	systemMessage := ""

	for {
		if stage.GetThisLocation().IsGoal {
			utils.PrintWIN()
			break
		}
		Print(&systemMessage, stage, char)
		Act(&char, &systemMessage, &stage)
	}
}

func PlayStage2(char unit.Character) {
	stage := maps.GetStage2()

	systemMessage := ""

	for {
		if stage.GetThisLocation().IsGoal {
			utils.PrintWIN()
			break
		}
		Print(&systemMessage, stage, char)
		Act(&char, &systemMessage, &stage)
	}
}
