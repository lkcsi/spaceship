package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lkcsi/spaceship/game"
)

var ship *game.Ship
var display *game.Display
var physics game.Physics

const (
	width  = 500
	height = 500
)

type Game struct {
}

func init() {
	ship = game.NewShip(width/2, height/2, 10, 20)
	display = game.NewDisplay(ship)
	physics = game.NewSpacePhysics()
}

func (g *Game) Update() error {
	var inputs = game.GetDirection(ebiten.KeyS, ebiten.KeyE, ebiten.KeyQ,
		ebiten.KeyD, ebiten.KeyA, ebiten.KeySpace)
	ship.UpdateShip(inputs)
	physics.UpdateShip(ship)

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
