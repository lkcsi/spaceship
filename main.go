package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lkcsi/spaceship/game"
)

var player *game.Ship

const (
	width  = 300
	height = 300
)

type Game struct {
}

func init() {
	player = game.NewShip(width/2, height/2)
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
