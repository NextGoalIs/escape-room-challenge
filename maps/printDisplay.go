package maps

import (
	"escape-room-challenge/rooms"
	"escape-room-challenge/utils"
	"fmt"
)

func PrintDisplay(systemMessage string, PADDING int, defaultMap [6][8]rooms.Room, nowX int, nowY int, myItemsString string, ableCommandsString string, directionRoom [4]*rooms.Room) {
	north := directionRoom[0]
	west := directionRoom[1]
	east := directionRoom[2]
	south := directionRoom[3]

	utils.ClearConsoleWindows()
	fmt.Println(systemMessage)
	println(utils.GetStringCenter(north.Name, PADDING-len(west.Name)))
	println(utils.GetStringCenter(west.Name+" "+defaultMap[nowX][nowY].Name+" "+east.Name, PADDING))
	println(utils.GetStringCenter(south.Name, PADDING-len(west.Name)))
	println()
	fmt.Println("가지고 있는 물건 : " + myItemsString)
	fmt.Println("행동 : " + ableCommandsString)

	print(">>>  ")
}
