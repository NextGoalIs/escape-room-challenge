package main

import (
	"bufio"
	"escape-room-challenge/maps"
	"escape-room-challenge/system"
	"escape-room-challenge/unit"
	"escape-room-challenge/utils"
	"os"
)

func main() {

	char := unit.Character{}
	char.SetName()

	stage := maps.GetStage1()

	systemMessage := ""
	isLookAtRoom := false

	for {
		var ableCommands []string
		reader := bufio.NewReader(os.Stdin)

		if stage.GetThisLocation().IsGoal {
			utils.PrintWIN()
			break
		}

		system.AddCommands(stage, &ableCommands, char)

		system.Print(isLookAtRoom, &systemMessage, stage, ableCommands)
		input := system.Scan(reader)

		//나머지 처리 미완
		// system.AddLookAtMessage(firstCommand)
		system.Act(input, ableCommands, &char, &systemMessage, &isLookAtRoom, &stage)
	}

}
