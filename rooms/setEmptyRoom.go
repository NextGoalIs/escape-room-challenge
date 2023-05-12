package rooms

import "escape-room-challenge/types"

func (r *Room) SetEmptyRoom() {
	r.DoorType = types.NoDoorType
	r.Icon = string(types.NoDoorIcon)
}
