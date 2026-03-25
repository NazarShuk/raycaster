package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player
	Walls []Wall
	Raycaster
}

var game = &Game{}

func (g *Game) Update() error {

	g.Player.update()
	g.Raycaster.update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.draw(screen)

	for i := 0; i < len(game.Walls); i++ {
		game.Walls[i].draw(screen)
	}

	g.Raycaster.draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game.Player.Size.X = 16
	game.Player.Size.Y = 16

	createWalls()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func createWalls() {
	game.Walls = append(game.Walls, Wall{
		Position: Vector2{
			X: 0,
			Y: 0,
		},
		Size: Vector2{
			X: 640,
			Y: 32,
		},
	})
	game.Walls = append(game.Walls, Wall{
		Position: Vector2{
			X: 0,
			Y: 0,
		},
		Size: Vector2{
			X: 32,
			Y: 480,
		},
	})

	game.Walls = append(game.Walls, Wall{
		Position: Vector2{
			X: 0,
			Y: 240 - 32,
		},
		Size: Vector2{
			X: 640,
			Y: 32,
		},
	})

	game.Walls = append(game.Walls, Wall{
		Position: Vector2{
			X: 320 - 32,
			Y: 0,
		},
		Size: Vector2{
			X: 32,
			Y: 480,
		},
	})

	game.Walls = append(game.Walls, Wall{
		Position: Vector2{
			X: 100,
			Y: 50,
		},
		Size: Vector2{
			X: 48,
			Y: 48,
		},
	})
}
