package game

type World struct {
	ships []*Ship
}

func NewWorld() *World {
	return &World{[]*Ship{}}
}

func (world *World) Update() {

}

func (world *World) Add(ship *Ship) {
	world.ships = append(world.ships, ship)
}
