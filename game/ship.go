package game

import (
	"slices"
)

type Ship struct {
	Pos      Vector
	Movement Vector
	Angle    float64
	Action   int
	Steering int
}

func (ship *Ship) UpdateShip(dirs []int) {
	ship.decel()
	ship.Action = Idle
	ship.Steering = Idle

	if slices.Contains(dirs, Throttle) {
		ship.Action = Throttle
	} else if slices.Contains(dirs, Stabilize) {
		ship.Action = Stabilize
	}
	if slices.Contains(dirs, Left) {
		ship.Steering = Left
	} else if slices.Contains(dirs, Right) {
		ship.Steering = Right
	}
}

func NewShip(x, y float64) *Ship {
	var ship = Ship{}
	ship.Movement = Vector{0, 0}
	ship.Pos = Vector{x, y}
	ship.Angle = 180
	ship.Steering = Idle
	ship.Action = Idle

	return &ship
}
