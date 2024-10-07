package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Display struct {
	ship       *Ship
	shipIm     *ebiten.Image
	throttleIm *ebiten.Image
	thWidth    int
	thHeight   int
}

func NewDisplay(ship *Ship) *Display {
	im := ebiten.NewImage(ship.Width, ship.Height)
	im.Fill(color.RGBA{255, 0, 0, 0})

	x := ship.Width / 2
	th_img := ebiten.NewImage(x, x)
	th_img.Fill(color.RGBA{255, 255, 255, 0})

	return &Display{ship: ship, shipIm: im, throttleIm: th_img, thWidth: x, thHeight: x}
}

func (d *Display) DrawShip(screen *ebiten.Image) {

	d.drawShip(screen)

	d.drawThrottle2(screen)

	if d.ship.Accel {
		d.drawThrottle(screen)
	}
}

func (d *Display) drawThrottle(screen *ebiten.Image) {
	var op = &ebiten.DrawImageOptions{}
	x, y := d.ship.Width/-4, -1*d.ship.Height/2-d.thHeight
	op.GeoM.Translate(float64(x), float64(y))
	d.drawRect(screen, d.throttleIm, op)
}

func (d *Display) drawThrottle2(screen *ebiten.Image) {
	var op = &ebiten.DrawImageOptions{}
	x, y := d.ship.Width/2, d.ship.Width/2
	op.GeoM.Translate(float64(x), float64(y))
	d.drawRect(screen, d.throttleIm, op)
}

func (d *Display) drawShip(screen *ebiten.Image) {
	var op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(d.ship.Width/-2), float64(d.ship.Height/-2))
	d.drawRect(screen, d.shipIm, op)
}

func (d *Display) drawRect(screen, image *ebiten.Image, op *ebiten.DrawImageOptions) {
	op.GeoM.Rotate((360 - d.ship.Angle) * math.Pi / 180)
	op.GeoM.Translate(d.ship.Pos.X, d.ship.Pos.Y)
	screen.DrawImage(image, op)
}
