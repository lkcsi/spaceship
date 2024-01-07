package game

const (
	MaxVel        = 3
	Accel         = 0.1
	Decel         = 0.01
	Fall          = 0.02
	SteeringAccel = 2.0
)

type World struct {
	ships []*Ship
}

func NewWorld() *World {
	return &World{[]*Ship{}}
}

func (world *World) Update() {
	world.updateShips()
}

func (world *World) Add(ship *Ship) {
	world.ships = append(world.ships, ship)
}

func (world *World) updateShips() {
	for _, ship := range world.ships {
		ship.decel()
		ship.steer()
		ship.accel()
		ship.move()
	}
}

func (ship *Ship) accel() {
	if ship.Action == Throttle {
		moveVector := Vector{Accel, Accel}
		moveVector.Rotate(ship.Angle)

		ship.Movement.Add(moveVector)
	} else if ship.Action == Stabilize {
		if ship.Angle > 180 {
			ship.Angle -= SteeringAccel
		} else if ship.Angle < 180 {
			ship.Angle += SteeringAccel
		}
	}
}

func (ship *Ship) steer() {
	if ship.Steering == Left {
		ship.Angle += (SteeringAccel)
	} else if ship.Steering == Right {
		ship.Angle -= (SteeringAccel)
	}
	ship.Angle = float64(rune(ship.Angle) % 360)
}

func (ship *Ship) decel() {
	ship.Movement.X -= Decel * ship.Movement.X
	ship.Movement.Y += Fall
}

func (ship *Ship) move() {
	ship.Pos.Add(ship.Movement)
}
