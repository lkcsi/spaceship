package game

import (
	"slices"
)

type Ship struct {
	Pos      Vector
	movement Vector
	Angle    float64
	Accel    bool
	Width    int
	Height   int
}

type Steer int

const (
	accel               = 0.1
	decel               = 0.01
	steeringAccel       = 4.0
	steerLeft     Steer = 1
	steerRight    Steer = -1
)

func (ship *Ship) UpdateShip(inputs []int) {
	ship.Accel = false
	ship.decel()

	if slices.Contains(inputs, Throttle) {
		ship.accel()
	}
	if slices.Contains(inputs, Left) {
		ship.steer(steerLeft)
	} else if slices.Contains(inputs, Right) {
		ship.steer(steerRight)
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
}

func (ship *Ship) steer(steer Steer) {
	ship.Angle += float64(steer) * steeringAccel
	ship.Angle = float64(rune(ship.Angle) % 360)
}

func (ship *Ship) decel() {
	ship.Accel = false
	ship.movement.X -= decel * ship.movement.X
	ship.movement.Y -= decel * ship.movement.Y
}

func (ship *Ship) move() {
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
	ship.Angle = 180
	ship.Accel = false
	ship.Height = height
	ship.Width = width

	return &ship
}
