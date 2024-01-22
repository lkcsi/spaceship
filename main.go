package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lkcsi/spaceship/game"
)

var ship_1 *game.Ship
var ship_2 *game.Ship
var world *game.World
var display *game.Display

const (
	width  = 300
	height = 300
)

type Game struct {
}

func init() {
	ship_1 = game.NewShip(width/4, height/2, 10, 20, color.RGBA{255, 0, 0, 0})
	ship_2 = game.NewShip(width/4*3, height/2, 10, 20, color.RGBA{0, 255, 0, 0})

	world = game.NewWorld(width, height)
	world.Add(ship_1)
	world.Add(ship_2)
	display = &game.Display{}
	display.AddShip(ship_1)
	display.AddShip(ship_2)
}

func (g *Game) Update() error {
	if ship_1.Alive {
		var inputs = game.GetDirection(ebiten.KeyW, ebiten.KeyD, ebiten.KeyA)
		ship_1.UpdateShip(inputs)
	}

	if ship_2.Alive {
		var inputs = game.GetDirection(ebiten.KeyArrowUp, ebiten.KeyArrowRight, ebiten.KeyArrowLeft)
		ship_2.UpdateShip(inputs)
	}

	world.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	display.Draw(screen)
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
