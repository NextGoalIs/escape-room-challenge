package rooms

import "escape-room-challenge/types"

func (r *Room) SetEmptyRoom() {
	r.DoorType = types.NoDoor
	r.Name = "🔳"
}
