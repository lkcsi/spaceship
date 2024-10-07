package game

import (
	"slices"
)

type Ship struct {
	Pos       Vector
	movement  Vector
	rotation  float64
	Angle     float64
	Accel     bool
	FrontLeft bool
	FrontRigt bool
	RearLeft  bool
	RearRight bool
	Width     int
	Height    int
}

type Steer int

const (
	accel               = 0.1
	decel               = 0.00
	steeringAccel       = 0.1
	steerLeft     Steer = 1
	steerRight    Steer = -1
)

func (ship *Ship) UpdateShip(inputs []int) {
	ship.Accel = false
	ship.FrontLeft = false
	ship.FrontRigt = false
	ship.RearLeft = false
	ship.RearRight = false

	ship.decel()

	ship.Angle = float64(rune(ship.Angle) % 360)

	if slices.Contains(inputs, Throttle) {
		ship.accel()
	}
	if slices.Contains(inputs, FrontLeft) {
		ship.rotate(steerLeft)
	} else if slices.Contains(inputs, FrontRight) {
		ship.rotate(steerRight)
	}
	if slices.Contains(inputs, Stabilize) {
		ship.stabilize()
	}

	ship.move()
}

func (ship *Ship) accel() {
	ship.Accel = true
	moveVector := Vector{accel, accel}
	moveVector.Rotate(ship.Angle)

	ship.movement.Add(moveVector)
}

func (ship *Ship) stabilize() {
	ship.movement = Vector{0, 0}
	ship.rotation = 0
}

func (ship *Ship) rotate(steer Steer) {
	ship.rotation += float64(steer) * steeringAccel
}

func (ship *Ship) decel() {
	ship.Accel = false
	ship.movement.X -= decel * ship.movement.X
	ship.movement.Y -= decel * ship.movement.Y

	ship.rotation -= decel * ship.rotation
}

func (ship *Ship) move() {
	ship.Angle += ship.rotation
	ship.Pos.Add(ship.movement)
}

func (ship *Ship) Front() *Vector {
	return ship.body(ship.Angle, float64(ship.Height/2))
}

func (ship *Ship) Rear() *Vector {
	return ship.body(ship.Angle, float64(ship.Height/-2))
}

func (ship *Ship) Left() *Vector {
	angle := rune(ship.Angle+90) % 360
	return ship.body(float64(angle), float64(ship.Width/2))
}

func (ship *Ship) Right() *Vector {
	angle := rune(ship.Angle+90) % 360
	return ship.body(float64(angle), float64(ship.Width/-2))
}

func (ship *Ship) body(angle, magnitude float64) *Vector {
	result := ship.Pos
	moveVector := &Vector{magnitude, magnitude}
	moveVector.Rotate(angle)
	result.Add(*moveVector)
	return &result
}

func NewShip(x, y float64, width, height int) *Ship {
	var ship = Ship{}
	ship.movement = Vector{0, 0}
	ship.Pos = Vector{x, y}
	ship.Angle = 0
	ship.Accel = false
	ship.Height = height
	ship.Width = width

	return &ship
}
