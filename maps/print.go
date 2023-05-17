package maps

import (
	"escape-room-challenge/rooms"
	"escape-room-challenge/utils"
	"fmt"
)

const PADDING int = 60

func Print(systemMessage string, defaultMap [6][8]rooms.Room, nowX int, nowY int, myItemsString string, ableCommandsString string, connectingRooms [4]*rooms.Room) {
	north := connectingRooms[0]
	west := connectingRooms[1]
	east := connectingRooms[2]
	south := connectingRooms[3]

	utils.ClearConsoleWindows()
	fmt.Println(systemMessage)
	println(utils.GetStringCenter(north.Icon, PADDING-len(west.Icon)))
	println(utils.GetStringCenter(west.Icon+" "+defaultMap[nowX][nowY].Icon+" "+east.Icon, PADDING))
	println(utils.GetStringCenter(south.Icon, PADDING-len(west.Icon)))
	println()
	fmt.Println("가지고 있는 물건 : " + myItemsString)
	fmt.Println("행동 : " + ableCommandsString)

	print(">>>  ")
}
