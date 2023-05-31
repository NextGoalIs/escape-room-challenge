package system

import (
	"escape-room-challenge/maps"
	"escape-room-challenge/unit"
	"escape-room-challenge/utils"
)

func PlayStage1(char unit.Character) {
	stage := maps.GetStage1()
	message := Message{}

	for {
		if stage.GetThisLocation().IsGoal {
			utils.PrintWIN()
			break
		}
		Print(&message, stage, char)
		Act(&char, &message, &stage)
	}
}

func PlayStage2(char unit.Character) {
	stage := maps.GetStage2()
	message := Message{}

	for {
		if stage.GetThisLocation().IsGoal {
			utils.PrintWIN()
			break
		}
		Print(&message, stage, char)
		Act(&char, &message, &stage)
	}
}
