package unit

import (
	"fmt"
	"math"
)

type Enemy struct {
	Name string

	Health       int
	AttackDamage int
	Defence      int
}

func (e Enemy) AttackTo(c *Character) {
	damage := c.GetCalcDefence() - e.AttackDamage
	if damage >= 0 {
		return
	}
	c.Health = c.Health + damage
	fmt.Println(math.Abs(float64(damage)), "만큼의 데미지를 입었습니다.")
}
func (e *Enemy) setHealth(h int) {
	e.Health = h
}

func (e Enemy) getHelath() int {
	return e.Health
}

func (e Enemy) getAttackDamage() int {
	return e.AttackDamage
}

func (e Enemy) getDefence() int {
	return e.Defence
}
