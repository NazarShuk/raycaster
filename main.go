package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player
}

func (g *Game) Update() error {

	g.Player.update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{}
	game.Player.Size.X = 16
	game.Player.Size.Y = 16

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
