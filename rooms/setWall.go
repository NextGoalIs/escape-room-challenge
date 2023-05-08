package rooms

import "escape-room-challenge/types"

func GetWall() *Room {
	room := NewRoom(types.NoDoorType, types.NoItem, false)
	room.Name = string(types.WallName)

	return &room
}
