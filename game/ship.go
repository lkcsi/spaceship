package game

import (
	"slices"
)

type Ship struct {
	Pos        Vector
	Movement   Vector
	Rotation   float64
	Angle      float64
	Accel      bool
	FrontLeft  bool
	FrontRight bool
	RearLeft   bool
	RearRight  bool
	Stabilize  bool
	Width      int
	Height     int
}

type Dir int

func (ship *Ship) UpdateShip(inputs []int) {
	ship.Accel = false
	ship.FrontLeft = false
	ship.FrontRight = false
	ship.RearLeft = false
	ship.RearRight = false
	ship.Stabilize = false

	// ship.decel()

	// ship.Angle = float64(rune(ship.Angle) % 360)

	if slices.Contains(inputs, Stabilize) {
		ship.Stabilize = true
		return
	}
	if slices.Contains(inputs, Throttle) {
		ship.Accel = true
	}
	if slices.Contains(inputs, FrontLeft) {
		ship.FrontLeft = true
	}
	if slices.Contains(inputs, FrontRight) {
		ship.FrontRight = true
	}
	if slices.Contains(inputs, RearLeft) {
		ship.RearLeft = true
	}
	if slices.Contains(inputs, RearRight) {
		ship.RearRight = true
	}
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
	ship.Movement = Vector{0, 0}
	ship.Accel = false
	ship.Rotation = 0
	ship.Pos = Vector{x, y}
	ship.Angle = 180
	ship.Height = height
	ship.Width = width

	return &ship
}
