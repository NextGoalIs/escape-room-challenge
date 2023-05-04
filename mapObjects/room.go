package mapObjects

type Room struct {
	Name string

	DoorType bool
	ItemType bool
	IsGoal   bool
}

func (r Room) GetName() string {
	return r.Name
}

func (r Room) SetName(name string) {
	r.Name = name
}

func NewRoom(doorType bool, itemType bool, isGoal bool) Room {

	room := Room{Name: "방"}

	if doorType {
		room.Name = "🚪"
		room.DoorType = doorType
	}

	if itemType {
		room.Name = "🔨"
		room.ItemType = itemType
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