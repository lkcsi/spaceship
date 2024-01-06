package game

import (
	"log"
	"slices"
)

const (
	MaxVel           = 3
	Accel            = 0.1
	Decel            = 0.01
	Fall             = 0.02
	SteeringAccel    = 2.0
	SteeringTrashold = 0.1
)

type Ship struct {
	Pos      Vector
	movement Vector
	Angle    float64
	velocity float64
}

func (ship *Ship) UpdateShip(dirs []int) {
	ship.decel()

	if slices.Contains(dirs, Throttle) {
		ship.accel()
	}
	if slices.Contains(dirs, Left) {
		ship.steer(1)
	} else if slices.Contains(dirs, Right) {
		ship.steer(-1)
	}

	ship.move()
}

func (ship *Ship) accel() {
	moveVector := Vector{Accel, Accel}
	moveVector.Rotate(ship.Angle)
	log.Println(moveVector)

	ship.movement.Add(moveVector)
}

func (ship *Ship) steer(dir int) {
	ship.Angle += float64(dir * SteeringAccel)

	if ship.Angle < 0 {
		ship.Angle = 360 + ship.Angle
	} else if ship.Angle > 360 {
		ship.Angle = float64(rune(ship.Angle) % 360)
	}
}

func (ship *Ship) decel() {

	ship.movement.X -= Decel * ship.movement.X
	ship.movement.Y += Fall
}

func (ship *Ship) move() {
	ship.Pos.Add(ship.movement)
}

func NewShip(x, y float64) *Ship {
	return &Ship{Vector{x, y}, Vector{0, 0}, 180, 0}
}
