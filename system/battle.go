package system

import (
	"escape-room-challenge/unit"
	"fmt"
	"math/rand"
	"time"
)

func Battle(c *unit.Character, e *unit.Enemy) {

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
		input := Scan()

		switch input {
		case "공격":
			c.AttackTo(e)
			e.AttackTo(c)
		case "도망":
			rand.New(rand.NewSource(time.Now().UnixNano()))
			switch rand.Intn(2) {
			case 1:
				//도망 성공
				GetMessageInstance().SetCompleteRun()
				return
			default:
				//도망 실패
				GetMessageInstance().SetFailRun()
				GetMessageInstance().Print()
				e.AttackTo(c)
			}
		default:
		}

	}
}
