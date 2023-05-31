package system

import (
	"escape-room-challenge/maps"
	"escape-room-challenge/unit"
	"escape-room-challenge/utils"
)

func PlayStage1(char unit.Character) {
	stage := maps.GetStage1()

	for {
		if stage.GetThisLocation().IsGoal {
			utils.PrintWIN()
			break
		}
		Print(stage, char)
		Act(&char, &stage)
	}
}

func PlayStage2(char unit.Character) {
	stage := maps.GetStage2()

	for {
		if stage.GetThisLocation().IsGoal {
			utils.PrintWIN()
			break
		}
		if char.Health <= 0 {
			utils.PrintYouDie()
			break
		}
		Print(stage, char)
		Act(&char, &stage)
	}
}
