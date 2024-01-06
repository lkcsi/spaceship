package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	Throttle = iota
	Down
	Left
	Right
	Idle
)

func GetDirection() []int {
	var dirs []int
	if inpututil.KeyPressDuration(ebiten.KeyArrowUp) > 0 {
		dirs = append(dirs, Throttle)
	}
	if inpututil.KeyPressDuration(ebiten.KeyArrowDown) > 0 {
		dirs = append(dirs, Down)
	}
	if inpututil.KeyPressDuration(ebiten.KeyArrowRight) > 0 {
		dirs = append(dirs, Right)
	}
	if inpututil.KeyPressDuration(ebiten.KeyArrowLeft) > 0 {
		dirs = append(dirs, Left)
	}
	return dirs
}
