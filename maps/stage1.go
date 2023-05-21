package maps

import (
	"escape-room-challenge/rooms"
	"escape-room-challenge/types"
)

func GetStage1() MapStruct {
	var defaultMap = setRooms(6, 8)
	defaultMap[3][0] = rooms.NewRoom(types.NoDoorType, types.Hammer, false, types.NoUnit)    // 망치
	defaultMap[1][1] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.Character) // 스타트지점
	defaultMap[2][1] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	defaultMap[3][1] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	defaultMap[3][2] = rooms.NewRoom(types.GlassType, types.NoItem, false, types.NoUnit) // 유리문 방
	defaultMap[1][3] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	defaultMap[2][3] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	defaultMap[3][3] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	defaultMap[4][3] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	defaultMap[0][4] = rooms.NewRoom(types.NoDoorType, types.Key, false, types.NoUnit)  // 열쇠
	defaultMap[1][4] = rooms.NewRoom(types.WoodType, types.NoItem, false, types.NoUnit) // 나무문
	defaultMap[4][4] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	defaultMap[5][4] = rooms.NewRoom(types.WoodType, types.NoItem, false, types.NoUnit) // 나무문
	defaultMap[5][5] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	defaultMap[5][6] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	defaultMap[5][7] = rooms.NewRoom(types.LockedType, types.NoItem, true, types.NoUnit) // 열쇠가 필요한 문

	defaultMap[1][1].SetMyCharacter()

	return MapStruct{nowX: 1, nowY: 1, nowMap: defaultMap}
}
