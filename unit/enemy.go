package unit

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Enemy struct {
	Name string

	Health       int
	AttackDamage int
	Defence      int

	dropTable map[string]int
}

func (e Enemy) AttackTo(c *Character) {
	damage := c.GetCalcDefence() - e.AttackDamage
	if damage >= 0 {
		return
	}
	c.Health = c.Health + damage
	fmt.Println(math.Abs(float64(damage)), "만큼의 데미지를 입었습니다.")
}

func (e Enemy) GiveItemTo(userItem []string) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	n := rand.Intn(100)
	for item, value := range e.dropTable {
		if n < value {
			userItem = append(userItem, item)
			if item == "회복약회복약" {
				userItem = append(userItem, item)
			}
			return item
		}
	}
	return ""
}
