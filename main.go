package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lkcsi/spaceship/game"
)

var ship_1 *game.Ship
var display *game.Display

const (
	width  = 500
	height = 500
)

type Game struct {
}

func init() {
	ship_1 = game.NewShip(width/4, height/2, 10, 20)
	display = game.NewDisplay(ship_1)
}

func (g *Game) Update() error {
	var inputs = game.GetDirection(ebiten.KeyW, ebiten.KeyD, ebiten.KeyA, ebiten.KeyControlLeft)
	ship_1.UpdateShip(inputs)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	display.DrawShip(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("Spaceship!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
