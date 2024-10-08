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

	if d.ship.FrontLeft {
		d.drawFrontLeft(screen)
	}
	if d.ship.FrontRight {
		d.drawFrontRight(screen)
	}
	if d.ship.RearRight {
		d.drawRearRight(screen)
	}
	if d.ship.RearLeft {
		d.drawRearLeft(screen)
	}
	if d.ship.Accel {
		d.drawThrottle(screen)
	}
}

func (d *Display) drawThrottle(screen *ebiten.Image) {
	x, y := -1*d.ship.Width/4, -1*d.ship.Height/2-d.thHeight
	d.drawBurn(screen, x, y)
}

func (d *Display) drawFrontRight(screen *ebiten.Image) {
	x, y := -1*d.ship.Width/2-d.thWidth, d.ship.Height/2-d.thHeight
	d.drawBurn(screen, x, y)
}

func (d *Display) drawFrontLeft(screen *ebiten.Image) {
	x, y := d.ship.Width/2, d.ship.Height/2-d.thHeight
	d.drawBurn(screen, x, y)
}

func (d *Display) drawRearRight(screen *ebiten.Image) {
	x, y := -1*d.ship.Width/2-d.thWidth, -1*d.ship.Height/2
	d.drawBurn(screen, x, y)
}

func (d *Display) drawRearLeft(screen *ebiten.Image) {
	x, y := d.ship.Width/2, -1*d.ship.Height/2
	d.drawBurn(screen, x, y)
}

func (d *Display) drawBurn(screen *ebiten.Image, x, y int) {
	var op = &ebiten.DrawImageOptions{}
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
