package unit

type Enemy struct {
	Name string

	Health       int
	AttackDamage int
	Defence      int
}

func (e Enemy) AttackTo(c *Character) {
	damage := c.Defence - e.AttackDamage
	if damage >= 0 {
		return
	}
	c.Health = c.Health + damage
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
