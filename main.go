package main

import (
	"bufio"
	"escape-room-challenge/maps"
	"escape-room-challenge/system"
	"escape-room-challenge/unit"
	"escape-room-challenge/utils"
	"os"
	"strings"
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
		var input string

		if stage.GetThisLocation().IsGoal {
			utils.PrintWIN()
			break
		}

		system.AddCommands(stage, &ableCommands, char)

		switch isLookAtRoom {
		case true:
			maps.LookAtRoomPrint(systemMessage, stage, ableCommands)
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
			// fmt.Scanln(&firstCommand, &secondCommand, &thirdCommand)
		default:
			maps.Print(systemMessage, stage, ableCommands)
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
			// fmt.Scanln(&firstCommand, &secondCommand, &thirdCommand)
		}

		//나머지 처리 미완
		// system.AddLookAtMessage(firstCommand)
		system.Act(input, ableCommands, &char, &systemMessage, &isLookAtRoom, &stage)
	}

}
