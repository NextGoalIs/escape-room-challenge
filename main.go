package main

import (
	"escape-room-challenge/character"
	"escape-room-challenge/maps"
	"escape-room-challenge/system"
	"escape-room-challenge/types"
	"escape-room-challenge/utils"
	"fmt"
	"strings"
)

func main() {

	stage1 := maps.GetStage1()
	char := character.Character{}

	stage1.GetThisLocation().SetMyCharacter()
	systemMessage := ""

	for {
		var ableCommands []string
		var inputItemCommand, inputCommand string

		if stage1.GetThisLocation().IsGoal {
			utils.PrintWIN()
			break
		}

		//무조건 먹게 하지 말고 줍게 하기...다음에...
		if stage1.GetThisLocation().ItemType != types.NoItem {
			stage1.GetThisLocation().PickUpItem(&systemMessage, &char)
		}

		connectingRooms := stage1.GetConnectingRooms()

		system.AddMoveCommands(connectingRooms, &ableCommands)
		system.AddUseItemCommands(char.Items, &ableCommands)
		system.AddOpenDoorCommands(connectingRooms, &ableCommands)

		ableCommandsString := strings.Join(ableCommands, ", ")
		myItemsString := strings.Join(char.Items, ", ")

		maps.Print(systemMessage, stage1, myItemsString, ableCommandsString, connectingRooms)
		fmt.Scanln(&inputItemCommand, &inputCommand)

		switch inputCommand {
		case "사용":
			if maps.UseItem(inputItemCommand, ableCommandsString, &char.Items, connectingRooms) {
				systemMessage = "아이템을 사용했습니다."
				continue
			}
		case "열기":
			if maps.OpenDoor(inputItemCommand, ableCommandsString, connectingRooms, &char.Items) {
				systemMessage = "문을 열었습니다."
				continue
			}
		default:
			switch inputItemCommand[0] {
			case "동"[0], "서"[0], "남"[0], "북"[0]:
				if maps.Move(inputItemCommand, ableCommandsString, &stage1) {
					systemMessage = ""
					continue
				}

				fallthrough
			default:
				systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
				continue
			}
		}
	}

}
