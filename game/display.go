package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Display struct {
	ships []ShipImage
}

type ShipImage struct {
	ship     *Ship
	body     *ebiten.Image
	throttle *ebiten.Image
}

func NewDisplay() *Display {
	return &Display{}
}

func (d *Display) AddShip(ship *Ship) {
	im := ebiten.NewImage(ship.Width, ship.Height)
	im.Fill(ship.Color)

	th_img := ebiten.NewImage(ship.Width/2, ship.Width/2)
	th_img.Fill(color.RGBA{255, 255, 0, 0})

	d.ships = append(d.ships, ShipImage{ship, im, th_img})
}

func (d *Display) RemoveShip(rship *Ship) {
	index := -1
	for idx, ship := range d.ships {
		if ship.ship == rship {
			index = idx
			break
		}
	}
	if index != -1 {
		d.ships = append(d.ships[:index], d.ships[index+1:]...)
	}
}

func drawShip(screen *ebiten.Image, shipImage *ShipImage) {
	ship := shipImage.ship
	var op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(ship.Width/-2), float64(ship.Height/-2))
	op.GeoM.Rotate((360 - ship.Angle) * math.Pi / 180)
	op.GeoM.Translate(ship.Pos.X, ship.Pos.Y)

	screen.DrawImage(shipImage.body, op)

	if ship.Accel {
		width, heigth := ship.Width/2, ship.Width/2
		var op2 = &ebiten.DrawImageOptions{}
		op2.GeoM.Translate(float64(width/-2), float64(ship.Height/-2-heigth))
		op2.GeoM.Rotate((360 - ship.Angle) * math.Pi / 180)
		op2.GeoM.Translate(ship.Pos.X, ship.Pos.Y)
		screen.DrawImage(shipImage.throttle, op2)
	}

}

func (d *Display) Draw(screen *ebiten.Image) {
	for _, s := range d.ships {
		drawShip(screen, &s)
	}
}
