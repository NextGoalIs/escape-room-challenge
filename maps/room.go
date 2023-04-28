package maps

import (
	"escape-room-challenge/utils"
	"fmt"
)

type RoomType interface {
	getName() string
}

type Room struct {
	North RoomType
	East  RoomType
	South RoomType
	West  RoomType

	Name string
}

func (room Room) getName() string {
    return room.Name
}

func (room Room) Display() {

    utils.ClearConsoleWindows()
    println(utils.PrintStringCenter(room.North.getName()))
    println(utils.PrintStringCenter("🚧  "+room.Name+"  🚪"))
    println(utils.PrintStringCenter("🚧"))
    println()
    fmt.Println("가지고 있는 물건 : ")
    fmt.Println("할 수 있는 일 : ")

    print(">>>  ")
    fmt.Scanln()
    
}