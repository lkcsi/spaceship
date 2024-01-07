package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lkcsi/spaceship/game"
)

var ship *game.Ship
var world *game.World

const (
	width  = 300
	height = 300
)

type Game struct {
}

func init() {
	ship = game.NewShip(width/2, height/2)
	world = game.NewWorld()
	world.Add(ship)
}

func (g *Game) Update() error {
	var inputs = game.GetDirection()
	ship.UpdateShip(inputs)
	world.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	game.DrawShip(screen, ship)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	log.Print("Started")

	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("Spaceship!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
