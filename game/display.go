package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Display struct {
	ships   []ShipImage
	paddles []PaddleImage
}

type ShipImage struct {
	image    *ebiten.Image
	throttle *ebiten.Image
	ship     *Ship
}

type PaddleImage struct {
	image  *ebiten.Image
	paddle *Paddle
}

func NewDisplay() *Display {
	return &Display{}
}

func (d *Display) AddPaddle(paddle *Paddle) {
	im := ebiten.NewImage(paddle.Width, paddle.Height)
	im.Fill(color.RGBA{255, 255, 255, 0})

	d.paddles = append(d.paddles, PaddleImage{im, paddle})
}

func (d *Display) AddShip(ship *Ship) {
	im := ebiten.NewImage(ship.Width, ship.Height)
	im.Fill(color.RGBA{255, 0, 0, 0})

	th := ebiten.NewImage(ship.Width/2, ship.Width/2)
	th.Fill(color.RGBA{255, 255, 0, 0})

	d.ships = append(d.ships, ShipImage{im, th, ship})
}

func drawPaddle(screen *ebiten.Image, paddleImage *PaddleImage) {
	p := paddleImage.paddle
	im := paddleImage.image

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Pos.X), float64(p.Pos.Y))
	screen.DrawImage(im, op)
}

func drawShip(screen *ebiten.Image, shipImage *ShipImage) {
	ship := shipImage.ship
	var op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(ship.Width/-2), float64(ship.Height/-2))
	op.GeoM.Rotate((360 - ship.Angle) * math.Pi / 180)
	op.GeoM.Translate(ship.Pos.X, ship.Pos.Y)

	screen.DrawImage(shipImage.image, op)

	width, heigth := ship.Width/2, ship.Width/2
	var op2 = &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(float64(width/-2), float64(ship.Height/-2-heigth))
	op2.GeoM.Rotate((360 - ship.Angle) * math.Pi / 180)
	op2.GeoM.Translate(ship.Pos.X, ship.Pos.Y)
	screen.DrawImage(shipImage.throttle, op2)

}

func (d *Display) Draw(screen *ebiten.Image) {
	for _, s := range d.ships {
		drawShip(screen, &s)
	}
	for _, p := range d.paddles {
		drawPaddle(screen, &p)
	}
}
