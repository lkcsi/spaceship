package game

type World struct {
	ships  []*Ship
	width  int
	height int
}

func NewWorld(width, height int) *World {
	return &World{[]*Ship{}, width, height}
}

func (world *World) Update() {
	for _, ship := range world.ships {
		world.check_wall(ship, ship.Front())
		world.check_wall(ship, ship.Rear())
		world.check_wall(ship, ship.Left())
		world.check_wall(ship, ship.Right())
	}
}

func (world *World) check_wall(ship *Ship, pos *Vector) {
	if int(pos.X) > world.width || int(pos.Y) > world.height || pos.X < 0 || pos.Y < 0 {
		ship.destroy()
	}
}

func (world *World) Add(ship *Ship) {
	world.ships = append(world.ships, ship)
}

func (world *World) Remove(rship *Ship) {
	index := -1
	for idx, ship := range world.ships {
		if ship == rship {
			index = idx
			break
		}
	}
	if index != -1 {
		world.ships = append(world.ships[:index], world.ships[index+1:]...)
	}
}
