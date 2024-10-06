package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Display struct {
	ship     *Ship
	body     *ebiten.Image
	throttle *ebiten.Image
}

func NewDisplay(ship *Ship) *Display {
	im := ebiten.NewImage(ship.Width, ship.Height)
	im.Fill(color.RGBA{255, 0, 0, 0})

	th_img := ebiten.NewImage(ship.Width/2, ship.Width/2)
	th_img.Fill(color.RGBA{255, 255, 0, 0})

	return &Display{ship: ship, body: im, throttle: th_img}
}

func (d *Display) DrawShip(screen *ebiten.Image) {
	ship := d.ship
	shipImage := d.body
	var op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(ship.Width/-2), float64(ship.Height/-2))
	op.GeoM.Rotate((360 - ship.Angle) * math.Pi / 180)
	op.GeoM.Translate(ship.Pos.X, ship.Pos.Y)

	screen.DrawImage(shipImage, op)

	if ship.Accel {
		d.drawThrottle(screen)
	}
}

func (d *Display) drawThrottle(screen *ebiten.Image) {
	width, heigth := d.ship.Width/2, d.ship.Width/2
	var op2 = &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(float64(width/-2), float64(d.ship.Height/-2-heigth))
	op2.GeoM.Rotate((360 - d.ship.Angle) * math.Pi / 180)
	op2.GeoM.Translate(d.ship.Pos.X, d.ship.Pos.Y)
	screen.DrawImage(d.throttle, op2)
}
