package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Walls    []Wall
	Entities []Entity

	FOV int

	MainRaycaster *Raycaster
}

var game = &Game{}

func (g *Game) Update() error {

	for _, entity := range game.Entities {
		entity.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, entity := range game.Entities {
		entity.Draw(screen)
	}

	for i := 0; i < len(game.Walls); i++ {
		//game.Walls[i].draw(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {

	ebiten.SetCursorMode(ebiten.CursorModeCaptured)
	createWalls()

	player := &Player{}

	player.Size.X = 16
	player.Size.Y = 16
	player.Position.X = 320 / 2
	player.Position.Y = 240 / 2

	spawnEntity(player)
	spawnEntity(&Chaser{})

	game.FOV = 360

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Game")
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

	for _, wall := range game.Walls {
		spawnEntity(&wall)
	}
}
