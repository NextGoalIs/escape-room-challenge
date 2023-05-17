package main

import (
	"escape-room-challenge/maps"
	"escape-room-challenge/rooms"
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

	for {

		systemMessage := ""

		thisRoom := &defaultMap[nowX][nowY]

		if thisRoom.IsGoal {
			utils.PrintWIN()
			break
		}
	CommandSwitch:
		systemMessage, myItems = thisRoom.PickUpItem(systemMessage, myItems)

		east, ableCommands, west, south, north := maps.MakeConsoleMap(&defaultMap, nowX, nowY)
		connectingRooms := [4]*rooms.Room{north, west, east, south}

		var inputItem, inputCommand string

		for _, v := range myItems {
			ableCommands = append(ableCommands, v+" 사용")
		}

		for _, v := range connectingRooms {
			if v.DoorType == types.NoDoorType {
				continue
			}

			if v.DoorType == types.WoodType {
				ableCommands = append(ableCommands, "나무문 열기")
				continue
			}

			if v.DoorType == types.GlassType {
				ableCommands = append(ableCommands, "유리문 열기")
				continue
			}

			if v.DoorType == types.LockedType {
				ableCommands = append(ableCommands, "잠긴문 열기")
				continue
			}
		}

		ableCommandsString := strings.Join(ableCommands, ", ")
		myItemsString := strings.Join(myItems, ", ")

		maps.Print(systemMessage, defaultMap, nowX, nowY, myItemsString, ableCommandsString, connectingRooms)
		fmt.Scanln(&inputItem, &inputCommand)

		var hasActed bool

		if inputCommand != "" && inputCommand == "사용" {
			hasActed = maps.UseItem(inputItem, ableCommandsString, nowX, nowY, &myItems, connectingRooms)

			if hasActed {
				systemMessage = "아이템을 사용했습니다."
				goto CommandSwitch
			}
		}

		if inputCommand != "" && inputCommand == "열기" {
			hasActed = maps.OpenDoor(inputItem, ableCommandsString, &defaultMap, nowX, nowY, connectingRooms, &myItems)

			if hasActed {
				systemMessage = "문을 열었습니다."
				goto CommandSwitch
			}
		}

		hasActed = maps.Move(inputItem, ableCommandsString, &defaultMap, &nowX, &nowY)

		if !hasActed {
			systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
			goto CommandSwitch
		}
	}

}
