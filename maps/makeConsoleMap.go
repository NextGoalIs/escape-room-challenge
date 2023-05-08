package maps

import "escape-room-challenge/rooms"

func MakeConsoleMap(defaultMap *[6][8]rooms.Room, nowX int, nowY int) (*rooms.Room, []string, *rooms.Room, *rooms.Room, *rooms.Room) {
	var north, east, south, west *rooms.Room
	var ableCommands []string

	if (nowY+1 < len(defaultMap[nowX])) && defaultMap[nowX][nowY+1].Name != "" {
		east = &defaultMap[nowX][nowY+1]
		ableCommands = append(ableCommands, "동")

	} else {
		wall := rooms.Room{Name: "🚧"}
		east = &wall
	}

	if (nowY-1 >= 0) && defaultMap[nowX][nowY-1].Name != "" {
		west = &defaultMap[nowX][nowY-1]
		ableCommands = append(ableCommands, "서")

	} else {
		wall := rooms.Room{Name: "🚧"}
		west = &wall
	}

	if (nowX-1 >= 0) && defaultMap[nowX-1][nowY].Name != "" {
		south = &defaultMap[nowX-1][nowY]
		ableCommands = append(ableCommands, "남")

	} else {
		wall := rooms.Room{Name: "🚧"}
		south = &wall
	}

	if (nowX+1 < len(defaultMap) && nowY < len(defaultMap[nowX+1])) && defaultMap[nowX+1][nowY].Name != "" {
		north = &defaultMap[nowX+1][nowY]
		ableCommands = append(ableCommands, "북")
	} else {
		wall := rooms.Room{Name: "🚧"}
		north = &wall
	}
	return east, ableCommands, west, south, north
}
