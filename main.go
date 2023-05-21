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
	isLookAtRoom := false

	for {
		if stage.GetThisLocation().IsGoal {
			utils.PrintWIN()
			break
		}

		system.Print(isLookAtRoom, &systemMessage, stage, char)
		input := system.Scan()

		//나머지 처리 미완
		// system.AddLookAtMessage(firstCommand)
		system.Act(input, &char, &systemMessage, &isLookAtRoom, &stage)
	}

}
