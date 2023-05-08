package maps

import "escape-room-challenge/rooms"

func MakeConsoleMap(defaultMap *[6][8]rooms.Room, nowX int, nowY int) (*rooms.Room, []string, *rooms.Room, *rooms.Room, *rooms.Room) {
	var north, east, south, west *rooms.Room
	var ableCommands []string

	if (nowY+1 < len(defaultMap[nowX])) && defaultMap[nowX][nowY+1].Name != "" {
		east = &defaultMap[nowX][nowY+1]
		ableCommands = append(ableCommands, "동")

	} else {
		east = rooms.GetWall()
	}

	if (nowY-1 >= 0) && defaultMap[nowX][nowY-1].Name != "" {
		west = &defaultMap[nowX][nowY-1]
		ableCommands = append(ableCommands, "서")

	} else {
		west = rooms.GetWall()
	}

	if (nowX-1 >= 0) && defaultMap[nowX-1][nowY].Name != "" {
		south = &defaultMap[nowX-1][nowY]
		ableCommands = append(ableCommands, "남")

	} else {
		south = rooms.GetWall()
	}

	if (nowX+1 < len(defaultMap) && nowY < len(defaultMap[nowX+1])) && defaultMap[nowX+1][nowY].Name != "" {
		north = &defaultMap[nowX+1][nowY]
		ableCommands = append(ableCommands, "북")
	} else {
		north = rooms.GetWall()
	}
	return east, ableCommands, west, south, north
}
