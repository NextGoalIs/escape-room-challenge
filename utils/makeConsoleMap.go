package utils

import "escape-room-challenge/mapObjects"

func MakeConsoleMap(defaultMap [6][8]mapObjects.Room, nowX int, nowY int, east mapObjects.Room, ableCommands []string, west mapObjects.Room, south mapObjects.Room, north mapObjects.Room) (mapObjects.Room, []string, mapObjects.Room, mapObjects.Room, mapObjects.Room) {
	if (nowY+1 < len(defaultMap[nowX])) && defaultMap[nowX][nowY+1].Name != "" {
		east = defaultMap[nowX][nowY+1]
		ableCommands = append(ableCommands, "동")

	} else {
		wall := mapObjects.Room{Name: "🚧"}
		east = wall
	}

	if (nowY-1 >= 0) && defaultMap[nowX][nowY-1].Name != "" {
		west = defaultMap[nowX][nowY-1]
		ableCommands = append(ableCommands, "서")

	} else {
		wall := mapObjects.Room{Name: "🚧"}
		west = wall
	}

	if (nowX-1 >= 0) && defaultMap[nowX-1][nowY].Name != "" {
		south = defaultMap[nowX-1][nowY]
		ableCommands = append(ableCommands, "남")

	} else {
		wall := mapObjects.Room{Name: "🚧"}
		south = wall
	}

	if (nowX+1 < len(defaultMap) && nowY < len(defaultMap[nowX+1])) && defaultMap[nowX+1][nowY].Name != "" {
		north = defaultMap[nowX+1][nowY]
		ableCommands = append(ableCommands, "북")
	} else {
		wall := mapObjects.Room{Name: "🚧"}
		north = wall
	}
	return east, ableCommands, west, south, north
}
