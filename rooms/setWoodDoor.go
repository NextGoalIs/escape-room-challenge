package rooms

import "escape-room-challenge/types"

func (room *Room) SetWoodDoor() {
	room.DoorType = types.WoodType
	room.Icon = string(types.WoodIcon)
}
