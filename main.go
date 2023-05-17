package main

import (
	"escape-room-challenge/maps"
	"escape-room-challenge/system"
	"escape-room-challenge/types"
	"escape-room-challenge/utils"
	"fmt"
	"strings"
)

func main() {

	defaultMap := maps.GetMap()
	nowX := 1
	nowY := 1
	var myItems []string

	defaultMap[nowX][nowY].SetMyCharacter()
	systemMessage := ""

	for {
		thisRoom := &defaultMap[nowX][nowY]
		var ableCommands []string

		if thisRoom.IsGoal {
			utils.PrintWIN()
			break
		}

		//무조건 먹게 하지 말고 줍게 하기...다음에...
		if thisRoom.ItemType != types.NoItem {
			thisRoom.PickUpItem(&systemMessage, &myItems)
		}

		connectingRooms := maps.GetConnectingRooms(&defaultMap, nowX, nowY)

		var inputItem, inputCommand string

		system.AddMoveCommands(connectingRooms, &ableCommands)
		system.AddUseItemCommands(myItems, &ableCommands)
		system.AddOpenDoorCommands(connectingRooms, &ableCommands)

		ableCommandsString := strings.Join(ableCommands, ", ")
		myItemsString := strings.Join(myItems, ", ")

		maps.Print(systemMessage, defaultMap, nowX, nowY, myItemsString, ableCommandsString, connectingRooms)
		fmt.Scanln(&inputItem, &inputCommand)

		switch inputCommand {
		case "사용":
			if maps.UseItem(inputItem, ableCommandsString, nowX, nowY, &myItems, connectingRooms) {
				systemMessage = "아이템을 사용했습니다."
				continue
			}
		case "열기":
			if maps.OpenDoor(inputItem, ableCommandsString, &defaultMap, nowX, nowY, connectingRooms, &myItems) {
				systemMessage = "문을 열었습니다."
				continue
			}
		default:
			switch inputItem[0] {
			case "동"[0], "서"[0], "남"[0], "북"[0]:
				if maps.Move(inputItem, ableCommandsString, &defaultMap, &nowX, &nowY) {
					systemMessage = ""
					continue
				}

				systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
			default:
				systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
				continue
			}
		}
	}

}
