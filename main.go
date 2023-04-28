package main

import (
	"escape-room-challenge/maps"
)

func main() {
	
	room := maps.Room{Name: "나", North: maps.Room{Name: "위쪽"} }

	
	for true {
		room.Display()
	}

}