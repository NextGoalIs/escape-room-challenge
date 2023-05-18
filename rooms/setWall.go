package rooms

import "escape-room-challenge/types"

func GetWall() Room {
	room := NewRoom(types.NoDoorType, types.NoItem, false)
	room.Icon = string(types.WallIcon)

	return room
}
