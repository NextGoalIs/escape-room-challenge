package maps

import "escape-room-challenge/rooms"

func setRooms(rows int, columns int) [][]rooms.Room {
	var defaultMap = make([][]rooms.Room, rows)
	for i := range defaultMap {
		defaultMap[i] = make([]rooms.Room, columns)
	}

	return defaultMap
}
