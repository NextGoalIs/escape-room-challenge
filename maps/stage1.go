package maps

import (
	"escape-room-challenge/rooms"
	"escape-room-challenge/types"
)

func GetStage1() [][]rooms.Room {
	var defaultMap = make([][]rooms.Room, 6)
	for i := range defaultMap {
		defaultMap[i] = make([]rooms.Room, 8)
	}
	defaultMap[3][0] = rooms.NewRoom(types.NoDoorType, types.Hammer, false) // 망치
	defaultMap[1][1] = rooms.NewRoom(types.NoDoorType, types.NoItem, false) // 스타트지점
	defaultMap[2][1] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[3][1] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[3][2] = rooms.NewRoom(types.GlassType, types.NoItem, false) // 유리문 방
	defaultMap[1][3] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[2][3] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[3][3] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[4][3] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[0][4] = rooms.NewRoom(types.NoDoorType, types.Key, false)  // 열쇠
	defaultMap[1][4] = rooms.NewRoom(types.WoodType, types.NoItem, false) // 나무문
	defaultMap[4][4] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[5][4] = rooms.NewRoom(types.WoodType, types.NoItem, false) // 나무문
	defaultMap[5][5] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[5][6] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[5][7] = rooms.NewRoom(types.LockedType, types.NoItem, true) // 열쇠가 필요한 문

	return defaultMap
}
