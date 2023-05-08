package main

import (
	"escape-room-challenge/mapObjects"
	"escape-room-challenge/maps"
	"escape-room-challenge/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {

	defaultMap := maps.GetMap()
	nowX := 1
	nowY := 1
	var myItems []string

	PADDING := 80

	defaultMap[nowX][nowY].Name = "🏃"

	for {

		systemMessage := ""

		_, ok := interface{}(defaultMap[nowX][nowY]).(mapObjects.Room)
		if !ok {
			fmt.Println("Error")
			break
		}

		thisRoom := &defaultMap[nowX][nowY]

		if thisRoom.IsGoal {
			utils.PrintWIN()
			break
		}

		systemMessage, myItems = utils.PickUpCurrentRoomItem(thisRoom, systemMessage, myItems)

		var north, east, south, west mapObjects.Room
		// var canMoverNorth, canMoveEast, canMoveSouth, canMoveWest bool

		var ableCommands []string

		east, ableCommands, west, south, north = utils.MakeConsoleMap(defaultMap, nowX, nowY, east, ableCommands, west, south, north)

		var selectedCommand string

		for _, v := range myItems {
			ableCommands = append(ableCommands, v+" 사용")
		}

		ableCommandsString := strings.Join(ableCommands, ", ")
		myItemsString := strings.Join(myItems, ", ")

	CommandSwitch:
		utils.ClearConsoleWindows()
		fmt.Println(systemMessage) //시스템
		println(utils.GetStringCenter(north.GetName(), PADDING-len(west.GetName())))
		println(utils.GetStringCenter(west.GetName()+" "+defaultMap[nowX][nowY].GetName()+" "+east.GetName(), PADDING))
		println(utils.GetStringCenter(south.GetName(), PADDING-len(west.GetName())))
		println()
		fmt.Println("가지고 있는 물건 : " + myItemsString)
		fmt.Println("할 수 있는 행동 : " + ableCommandsString)

		print(">>>  ")
		fmt.Scan(&selectedCommand)

		hasActed := Move(selectedCommand, ableCommandsString, &defaultMap, &nowX, &nowY)

		if !hasActed {
			systemMessage = "할 수 없는 행동입니다."
			goto CommandSwitch
		}

		// if strings.Contains(ableCommandsString, selectedCommand) {
		// 	switch selectedCommand {
		// 	//이동 커맨드
		// 	case "북":
		// 		SetRoomNameDefaultRoom(&defaultMap, nowX, nowY)
		// 		nowX += 1
		// 		SetRoomNameMyIcon(&defaultMap, nowX, nowY)
		// 	case "동":
		// 		SetRoomNameDefaultRoom(&defaultMap, nowX, nowY)
		// 		nowY += 1
		// 		SetRoomNameMyIcon(&defaultMap, nowX, nowY)
		// 	case "남":
		// 		SetRoomNameDefaultRoom(&defaultMap, nowX, nowY)
		// 		nowX -= 1
		// 		SetRoomNameMyIcon(&defaultMap, nowX, nowY)
		// 	case "서":
		// 		SetRoomNameDefaultRoom(&defaultMap, nowX, nowY)
		// 		nowY -= 1
		// 		SetRoomNameMyIcon(&defaultMap, nowX, nowY)
		// 	default:
		// 		fmt.Println("디폴트로 들어와버렸음")
		// 	}
		// } else {
		// 	systemMessage = "할 수 없는 행동입니다."
		// 	goto CommandSwitch
		// }

	}

}

func Move(selectedCommand string, ableCommandsString string, defaultMap *[6][8]mapObjects.Room, nowX *int, nowY *int) bool {
	reg, _ := regexp.Compile("^동.*")

	isEast := reg.MatchString(selectedCommand)

	if isEast && strings.Contains(ableCommandsString, "동") {
		SetRoomNameDefaultRoom(defaultMap, *nowX, *nowY)
		*nowY += 1
		SetRoomNameMyIcon(defaultMap, *nowX, *nowY)
		return true
	}

	reg, _ = regexp.Compile("^서.*")

	isWest := reg.MatchString(selectedCommand)

	if isWest && strings.Contains(ableCommandsString, "서") {
		SetRoomNameDefaultRoom(defaultMap, *nowX, *nowY)
		*nowY -= 1
		SetRoomNameMyIcon(defaultMap, *nowX, *nowY)
		return true

	}

	reg, _ = regexp.Compile("^남.*")

	isSouth := reg.MatchString(selectedCommand)

	if isSouth && strings.Contains(ableCommandsString, "남") {
		SetRoomNameDefaultRoom(defaultMap, *nowX, *nowY)
		*nowX -= 1
		SetRoomNameMyIcon(defaultMap, *nowX, *nowY)
		return true

	}

	reg, _ = regexp.Compile("^북.*")

	isNorth := reg.MatchString(selectedCommand)

	if isNorth && strings.Contains(ableCommandsString, "북") {
		SetRoomNameDefaultRoom(defaultMap, *nowX, *nowY)
		*nowX += 1
		SetRoomNameMyIcon(defaultMap, *nowX, *nowY)
		return true
	}

	return false
}

func SetRoomNameMyIcon(defaultMap *[6][8]mapObjects.Room, nowX int, nowY int) {
	defaultMap[nowX][nowY].Name = "🏃"
}

func SetRoomNameDefaultRoom(defaultMap *[6][8]mapObjects.Room, nowX int, nowY int) {
	defaultMap[nowX][nowY].Name = "🔳"
}
