package maps

import "escape-room-challenge/mapObjects"

func GetMap() [6][8]mapObjects.Room {
	var defaultMap [6][8]mapObjects.Room = [6][8]mapObjects.Room{}
	defaultMap[3][0] = mapObjects.NewRoom(0,2,false) // 망치
	defaultMap[1][1] = mapObjects.NewRoom(0,0,false) // 스타트지점
	defaultMap[2][1] = mapObjects.NewRoom(0,0,false)
	defaultMap[3][1] = mapObjects.NewRoom(0,0,false)
	defaultMap[3][2] = mapObjects.NewRoom(2,0,false) // 유리문 방
	defaultMap[1][3] = mapObjects.NewRoom(0,0,false)
	defaultMap[2][3] = mapObjects.NewRoom(0,0,false)
	defaultMap[3][3] = mapObjects.NewRoom(0,0,false)
	defaultMap[4][3] = mapObjects.NewRoom(0,0,false)
	defaultMap[0][4] = mapObjects.NewRoom(0,1,false) // 열쇠
	defaultMap[1][4] = mapObjects.NewRoom(1,0,false) // 나무문
	defaultMap[4][4] = mapObjects.NewRoom(0,0,false) 
	defaultMap[5][4] = mapObjects.NewRoom(1,0,false) // 나무문
	defaultMap[5][5] = mapObjects.NewRoom(0,0,false)
	defaultMap[5][6] = mapObjects.NewRoom(0,0,false)
	defaultMap[5][7] = mapObjects.NewRoom(3,0,true) // 열쇠가 필요한 문

	return defaultMap
}