package rooms

import "escape-room-challenge/types"

func GetWall() *Room {
	room := NewRoom(types.NoDoor, types.NoItem, false)
	room.Name = "🚧"

	return &room
}
