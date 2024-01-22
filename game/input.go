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

func GetDirection(up, right, left ebiten.Key) []int {
	var dirs []int
	if inpututil.KeyPressDuration(up) > 0 {
		dirs = append(dirs, Throttle)
	}
	if inpututil.KeyPressDuration(right) > 0 {
		dirs = append(dirs, Right)
	}
	if inpututil.KeyPressDuration(left) > 0 {
		dirs = append(dirs, Left)
	}
	//if inpututil.KeyPressDuration(ebiten.KeySpace) > 0 {
	//	dirs = append(dirs, Stabilize)
	//}
	return dirs
}
