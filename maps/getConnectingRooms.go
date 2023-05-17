package maps

import "escape-room-challenge/rooms"

func GetConnectingRooms(defaultMap *[6][8]rooms.Room, nowX int, nowY int) [4]*rooms.Room {
	var north, east, south, west *rooms.Room

	if (nowY+1 < len(defaultMap[nowX])) && defaultMap[nowX][nowY+1].Icon != "" {
		east = &defaultMap[nowX][nowY+1]

	} else {
		east = rooms.GetWall()
	}

	if (nowY-1 >= 0) && defaultMap[nowX][nowY-1].Icon != "" {
		west = &defaultMap[nowX][nowY-1]

	} else {
		west = rooms.GetWall()
	}

	if (nowX-1 >= 0) && defaultMap[nowX-1][nowY].Icon != "" {
		south = &defaultMap[nowX-1][nowY]

	} else {
		south = rooms.GetWall()
	}

	if (nowX+1 < len(defaultMap) && nowY < len(defaultMap[nowX+1])) && defaultMap[nowX+1][nowY].Icon != "" {
		north = &defaultMap[nowX+1][nowY]
	} else {
		north = rooms.GetWall()
	}
	return [4]*rooms.Room{east, west, south, north}
}
