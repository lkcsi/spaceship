package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var shipImage *ebiten.Image
var throttleImage *ebiten.Image

func init() {

	shipImage = ebiten.NewImage(10, 20)
	throttleImage = ebiten.NewImage(5, 5)

	shipImage.Fill(color.RGBA{255, 0, 0, 0})
	throttleImage.Fill(color.RGBA{255, 255, 0, 0})
}

func translateImage(ship *Ship, x, y float64) *ebiten.DrawImageOptions {
	var op1 = &ebiten.DrawImageOptions{}
	op1.GeoM.Translate(x, y)
	op1.GeoM.Rotate((360 - ship.Angle) * math.Pi / 180)
	op1.GeoM.Translate(ship.Pos.X, ship.Pos.Y)
	return op1
}

func DrawShip(screen *ebiten.Image, ship *Ship) {
	screen.DrawImage(shipImage, translateImage(ship, -5, -10))

	if ship.Action == Throttle {
		screen.DrawImage(throttleImage, translateImage(ship, -2.5, -15))
	}

}
