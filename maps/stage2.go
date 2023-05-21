package maps

import (
	"escape-room-challenge/rooms"
	"escape-room-challenge/types"
)

func GetStage2() mapStruct {
	var m = setRooms(10, 10)

	m[7][0] = rooms.NewRoom(types.LockedType, types.NoItem, true, types.NoUnit) // 끝문
	m[7][1] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	m[6][1] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	m[5][1] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	m[4][1] = rooms.NewRoom(types.GlassType, types.NoItem, false, types.NoUnit)
	m[3][1] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.Deer)
	m[3][2] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	m[3][3] = rooms.NewRoom(types.GlassType, types.NoItem, false, types.NoUnit)
	m[3][4] = rooms.NewRoom(types.NoDoorType, types.Hammer, false, types.Rabbit)
	m[4][4] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	m[5][4] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	m[6][4] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.Squirrel)
	m[6][3] = rooms.NewRoom(types.WoodType, types.Chest, false, types.NoUnit)
	m[7][4] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	m[8][4] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	// m[9][4] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit) 아이템이 두개라니...
	m[6][5] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	m[6][6] = rooms.NewRoom(types.NoDoorType, types.Potion, false, types.NoUnit)
	m[5][6] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.Squirrel)
	m[4][6] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	m[3][6] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	m[2][6] = rooms.NewRoom(types.NoDoorType, types.Hammer, false, types.NoUnit)
	m[1][6] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	m[0][6] = rooms.NewRoom(types.WoodType, types.Chest, false, types.NoUnit)
	m[6][7] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	m[2][7] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	m[6][8] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	m[2][8] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	m[6][9] = rooms.NewRoom(types.GlassType, types.Chest, false, types.NoUnit)
	m[2][9] = rooms.NewRoom(types.NoDoorType, types.WoodSword, false, types.NoUnit)
	m[1][9] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.NoUnit)
	m[0][9] = rooms.NewRoom(types.NoDoorType, types.NoItem, false, types.Character)

	m[0][9].SetMyCharacter()

	return mapStruct{nowX: 0, nowY: 9, nowMap: m}
}
