package main

import (
	"escape-room-challenge/maps"
	"escape-room-challenge/system"
	"escape-room-challenge/unit"
	"escape-room-challenge/utils"
	"fmt"
	"strings"
)

func main() {

	char := unit.Character{}
	char.SetName()

	stage1 := maps.GetStage1()

	systemMessage := ""
	isLookAtRoom := false

	for {
		var ableCommands []string
		var firstCommand, secondCommand, thirdCommand string

		if stage1.GetThisLocation().IsGoal {
			utils.PrintWIN()
			break
		}

		connectingRooms := stage1.GetConnectingRooms()

		system.AddMoveCommands(connectingRooms, &ableCommands)
		system.AddUseItemCommands(char.Items, &ableCommands)
		system.AddOpenDoorCommands(connectingRooms, &ableCommands)
		system.AddDefaultCommands(&ableCommands)

		myItemsString := strings.Join(char.Items, ", ")

		switch isLookAtRoom {
		case true:
			maps.LookAtRoomPrint(systemMessage, stage1, myItemsString, ableCommands, connectingRooms)
			fmt.Scanln(&firstCommand, &secondCommand, &thirdCommand)
		default:
			maps.Print(systemMessage, stage1, myItemsString, ableCommands, connectingRooms)
			fmt.Scanln(&firstCommand, &secondCommand, &thirdCommand)
		}

		switch secondCommand {
		case "사용":
			if maps.UseItem(firstCommand, ableCommands, &char.Items, connectingRooms, thirdCommand) {
				systemMessage = "아이템을 사용했습니다."
				isLookAtRoom = false
				continue
			}
		case "열기":
			if maps.OpenDoor(firstCommand, ableCommands, connectingRooms, &char.Items) {
				systemMessage = "문을 열었습니다."
				isLookAtRoom = false
				continue
			}
		case "보기":
			switch firstCommand {
			case "방":
				isLookAtRoom = true
				systemMessage = ""
				continue
			default:
				//나머지 처리 미완
				// system.AddLookAtMessage(firstCommand)
			}
			systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
			continue
		case "줍기":
			if stage1.GetThisLocation().PickUpItem(&systemMessage, &char) {
				isLookAtRoom = false
				continue
			}
			systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
			continue
		default:
			switch firstCommand[0] {
			case "동"[0], "서"[0], "남"[0], "북"[0]:
				if maps.Move(firstCommand, ableCommands, &stage1) {
					isLookAtRoom = false
					systemMessage = ""
					continue
				}

				fallthrough
			default:
				isLookAtRoom = false
				systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
				continue
			}
		}
	}

}
