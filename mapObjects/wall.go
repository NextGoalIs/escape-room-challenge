package mapObjects

type wall struct {
	name string
}

func (w wall) getName() string {
	return w.name
}

func NewWall() wall {
	return wall{name: "🚧"}
}