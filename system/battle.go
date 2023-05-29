package system

import (
	"escape-room-challenge/unit"
	"fmt"
)

func Battle(c unit.Character, e unit.Enemy) {

	for {
		if c.Health <= 0 {
			fmt.Println("패배")
			return
		}
		if e.Health <= 0 {
			fmt.Println("승리")
			return
		}
		fmt.Println("명령어 : 공격, 도망")
		fmt.Print(">>> ")
		Scan()
	}
}
