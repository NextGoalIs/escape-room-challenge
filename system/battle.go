package system

import (
	"escape-room-challenge/types"
	"escape-room-challenge/unit"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Battle(c *unit.Character, e *unit.Enemy) {

	for {
		if c.Health <= 0 {
			return
		}
		if e.Health <= 0 {
			GetMessageInstance().SetYouWin(e.Name)
			return
		}
		fmt.Println(string(types.MyCharacterIcon), strings.Repeat(" ", 30), string(types.EnemyIcon))
		fmt.Println("체력 : ", c.Health, strings.Repeat(" ", 30), "체력 : ", e.Health)
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
