package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	Throttle = iota
	FrontLeft
	FrontRight
	RearLeft
	RearRight
	Shoot
	Stabilize
	Idle
)

func GetDirection(throttle, fr, fl, rr, rl, stabilize ebiten.Key) []int {
	var dirs []int
	if inpututil.KeyPressDuration(throttle) > 0 {
		dirs = append(dirs, Throttle)
	}
	if inpututil.KeyPressDuration(fr) > 0 {
		dirs = append(dirs, FrontRight)
	}
	if inpututil.KeyPressDuration(fl) > 0 {
		dirs = append(dirs, FrontLeft)
	}
	if inpututil.KeyPressDuration(rr) > 0 {
		dirs = append(dirs, RearRight)
	}
	if inpututil.KeyPressDuration(rl) > 0 {
		dirs = append(dirs, RearLeft)
	}
	if inpututil.IsKeyJustPressed(stabilize) {
		dirs = append(dirs, Stabilize)
	}
	return dirs
}
