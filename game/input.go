package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	Throttle = iota
	Left
	Right
	Stabilize
	Idle
)

func GetDirection() []int {
	var dirs []int
	if inpututil.KeyPressDuration(ebiten.KeyArrowUp) > 0 {
		dirs = append(dirs, Throttle)
	}
	if inpututil.KeyPressDuration(ebiten.KeyArrowRight) > 0 {
		dirs = append(dirs, Right)
	}
	if inpututil.KeyPressDuration(ebiten.KeyArrowLeft) > 0 {
		dirs = append(dirs, Left)
	}
	if inpututil.KeyPressDuration(ebiten.KeySpace) > 0 {
		dirs = append(dirs, Stabilize)
	}
	return dirs
}
