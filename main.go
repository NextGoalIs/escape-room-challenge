package main

import (
	"escape-room-challenge/maps"
	"escape-room-challenge/system"
	"escape-room-challenge/unit"
	"escape-room-challenge/utils"
	"fmt"
)

func main() {

	char := unit.Character{}
	char.SetName()

	stage := maps.GetStage1()

	systemMessage := ""
	isLookAtRoom := false

	for {
		var ableCommands []string
		var firstCommand, secondCommand, thirdCommand string

		if stage.GetThisLocation().IsGoal {
			utils.PrintWIN()
			break
		}

		connectingRooms := stage.GetConnectingRooms()

		system.AddCommands(connectingRooms, &ableCommands, char)

		switch isLookAtRoom {
		case true:
			maps.LookAtRoomPrint(systemMessage, stage, ableCommands, connectingRooms)
			fmt.Scanln(&firstCommand, &secondCommand, &thirdCommand)
		default:
			maps.Print(systemMessage, stage, ableCommands, connectingRooms)
			fmt.Scanln(&firstCommand, &secondCommand, &thirdCommand)
		}

		//나머지 처리 미완
		// system.AddLookAtMessage(firstCommand)
		system.Act(firstCommand, secondCommand, thirdCommand, ableCommands, &char, connectingRooms, &systemMessage, &isLookAtRoom, &stage)
	}

}
