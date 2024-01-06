package game

import (
	"math"
)

type Vector struct {
	X float64
	Y float64
}

func (vector *Vector) Rotate(angle float64) {
	rad := angle * math.Pi / 180
	vector.X *= math.Sin(rad)
	vector.Y *= math.Cos(rad)
}

func (vector *Vector) Add(other Vector) {
	vector.X += other.X
	vector.Y += other.Y
}
