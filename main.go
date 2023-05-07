package main

import (
	"escape-room-challenge/mapObjects"
	"escape-room-challenge/maps"
	"escape-room-challenge/utils"
	"fmt"
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

		_,ok := interface{}(defaultMap[nowX][nowY]).(mapObjects.Room)
		if !ok {
			fmt.Println("Error")
			break
		} 

		thisRoom := defaultMap[nowX][nowY]

		if thisRoom.IsGoal {
			printHelloWorld()
			break
		}

		switch thisRoom.ItemType {
		case 0:
		case 1:
			systemMessage = "열쇠를 획득했다."
		thisRoom.ItemType = 0
			myItems = append(myItems, "열쇠")
		case 2:
			systemMessage = "망치를 획득했다."
		thisRoom.ItemType = 0
			myItems = append(myItems, "망치")
		default:
		}

		var north, east, south, west mapObjects.Room
		// var canMoverNorth, canMoveEast, canMoveSouth, canMoveWest bool

		var ableCommands []string

		if (nowY+1 < len(defaultMap[nowX])) && defaultMap[nowX][nowY+1].Name != "" {
			east = defaultMap[nowX][nowY+1]
			// canMoveEast = true
			ableCommands = append(ableCommands, "동")

		} else {
			wall := mapObjects.Room{Name: "🚧"}
			east = wall
		}

		if (nowY-1 >= 0) && defaultMap[nowX][nowY-1].Name != "" {
			west = defaultMap[nowX][nowY-1]
			// canMoveWest = true
			ableCommands = append(ableCommands, "서")

		} else {
			wall := mapObjects.Room{Name: "🚧"}
			west = wall
		}

		if (nowX-1 >= 0) && defaultMap[nowX-1][nowY].Name != "" {
			south = defaultMap[nowX-1][nowY]
			// canMoveSouth = true
			ableCommands = append(ableCommands, "남")

		} else {
			wall := mapObjects.Room{Name: "🚧"}
			south = wall
		}

		if (nowX+1 < len(defaultMap) && nowY < len(defaultMap[nowX+1])) && defaultMap[nowX+1][nowY].Name != "" {
			north = defaultMap[nowX+1][nowY]
			// canMoverNorth = true
			ableCommands = append(ableCommands, "북")
		} else {
			wall := mapObjects.Room{Name: "🚧"}
			north = wall
		}

	var selectedCommand string

	ableCommandsString := strings.Join(ableCommands, ", ")
	myItemsString := strings.Join(myItems, ", ")


	CommandSwitch:utils.ClearConsoleWindows()
	fmt.Println(systemMessage) //시스템 
	println(utils.GetStringCenter(north.GetName(), PADDING-len(west.GetName())))
	println(utils.GetStringCenter(west.GetName()+ " " + defaultMap[nowX][nowY].GetName() + " " + east.GetName(),PADDING))
	println(utils.GetStringCenter(south.GetName(),PADDING-len(west.GetName())))
	println()
	fmt.Println("가지고 있는 물건 : " + myItemsString)
	fmt.Println("할 수 있는 행동 : " + ableCommandsString )

	print(">>>  ")
	fmt.Scanln(&selectedCommand)

	if strings.Contains(ableCommandsString, selectedCommand) {
		switch selectedCommand {
			//이동 커맨드
		case "북":
			defaultMap[nowX][nowY].Name = "방"
			nowX += 1
			defaultMap[nowX][nowY].Name = "🏃"
		case "동":
			defaultMap[nowX][nowY].Name = "방"
			nowY += 1
			defaultMap[nowX][nowY].Name = "🏃"
		case "남":
			defaultMap[nowX][nowY].Name = "방"
			nowX -= 1
			defaultMap[nowX][nowY].Name = "🏃"
		case "서":
			defaultMap[nowX][nowY].Name = "방"
			nowY -= 1
			defaultMap[nowX][nowY].Name = "🏃"
			//문 열기
		
		default:
			fmt.Println("디폴트로 들어와버렸음")
		}
	} else {
		systemMessage = "할 수 없는 행동입니다."
		goto CommandSwitch
	}
	
	}

}

func printHelloWorld() {
	fmt.Println(`_    _  _____  _   _ 
| |  | ||_   _|| \ | |
| |  | |  | |  |  \| |
| |/\| |  | |  | .   |
\  /\  / _| |_ | |\  |
\/  \/  \___/ \_| \_/`)
}