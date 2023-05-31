package unit

func NewDear() Enemy {
	e := Enemy{Name: "사슴", Health: 100, AttackDamage: 10, Defence: 5, dropTable: map[string]int{"열쇠": 100}}

	return e
}
