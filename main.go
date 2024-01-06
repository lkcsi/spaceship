package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lkcsi/spaceship/game"
)

var player *game.Ship

type Game struct {
}

func init() {
	player = game.NewShip(160, 120)
}

func (g *Game) Update() error {
	var inputs = game.GetDirection()
	player.UpdateShip(inputs)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	game.DrawShip(screen, player)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	log.Print("Started")

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
