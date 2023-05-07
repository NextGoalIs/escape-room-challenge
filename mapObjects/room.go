package mapObjects

import "escape-room-challenge/types"

type Room struct {
	Name string

	DoorType types.DoorTypes
	ItemType types.ItemTypes
	IsGoal   bool
}

func (r Room) GetName() string {
	return r.Name
}

func (r Room) SetName(name string) {
	r.Name = name
}

func NewRoom(doorType types.DoorTypes, itemType types.ItemTypes, isGoal bool) Room {

	room := Room{Name: "방"}

	switch doorType {
	case 0:
	case 1:
		room.Name = "🚪"
		room.DoorType = doorType
	case 2:
		room.Name = "🧊"
		room.DoorType = doorType
	case 3 :
		room.Name = "🔒"
		room.DoorType = doorType
	default:
	}

	// if doorType {
	// 	room.Name = "🚪"
	// 	room.DoorType = doorType
	// }

	switch itemType {
	case 0:
	case 1:
		room.Name = "🔑"
		room.ItemType = itemType
	case 2:
		room.Name = "🔨"
		room.ItemType = itemType
	default:
	}

	

	if isGoal {
		room.IsGoal = isGoal
	}

	return room
}

// func (r Room) Display(m interface{}, nowX int, nowY int) {

// 	utils.ClearConsoleWindows()
// 	println(utils.GetStringCenter("북"))
// 	println(utils.GetStringCenter("🚧  " + r.GetName() + "  🚪"))
// 	println(utils.GetStringCenter("🚧"))
// 	println()
// 	fmt.Println("가지고 있는 물건 : ")
// 	fmt.Println("할 수 있는 일 : ")

// 	print(">>>  ")
// 	fmt.Scanln()

// }