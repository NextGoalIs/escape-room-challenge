package rooms

import "escape-room-challenge/types"

func GetWall() *Room {
	room := Room{IsWall: true}
	room.Icon = string(types.WallIcon)

	return &room
}
