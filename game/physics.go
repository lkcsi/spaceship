package game

type Physics interface {
	UpdateShip(*Ship)
}

const (
	accelSpeed        = 0.05
	decelSpeed        = 0.015
	steeringSpeed     = 0.05
	pos           Dir = 1
	neg           Dir = -1
)

type SpacePhyisics struct {
}

func NewSpacePhysics() *SpacePhyisics {
	return &SpacePhyisics{}
}

func (p *SpacePhyisics) UpdateShip(ship *Ship) {

	decel(ship)

	if ship.Stabilize {
		stabilize(ship)
		return
	}
	if ship.Accel {
		accel(ship)
	}
	if ship.FrontLeft && !ship.RearLeft {
		rotate(ship, neg)
	}
	if ship.FrontRight && !ship.RearRight {
		rotate(ship, pos)
	}
	if ship.RearLeft && !ship.FrontLeft {
		rotate(ship, pos)
	}
	if ship.RearRight && !ship.FrontRight {
		rotate(ship, neg)
	}
	if ship.RearRight && ship.FrontRight {
		shift(ship, pos)
	}
	if ship.RearLeft && ship.FrontLeft {
		shift(ship, neg)
	}
	move(ship)
}

func accel(ship *Ship) {
	moveVector := Vector{accelSpeed, accelSpeed}
	moveVector.Rotate(ship.Angle)

	ship.Movement.Add(moveVector)
}

func shift(ship *Ship, dir Dir) {
	moveVector := Vector{accelSpeed, accelSpeed}
	moveVector.Rotate(ship.Angle + (float64(dir) * 90))

	ship.Movement.Add(moveVector)
}

func stabilize(ship *Ship) {
	ship.Movement = Vector{0, 0}
	ship.Rotation = 0
}

func rotate(ship *Ship, steer Dir) {
	ship.Rotation += float64(steer) * steeringSpeed
}

func decel(ship *Ship) {
	ship.Movement.X -= decelSpeed * ship.Movement.X
	ship.Movement.Y -= decelSpeed * ship.Movement.Y

	ship.Rotation -= decelSpeed * ship.Rotation
}

func move(ship *Ship) {
	ship.Angle += ship.Rotation
	ship.Pos.Add(ship.Movement)
}
