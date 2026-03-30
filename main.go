package main

import (
	"cmp"
	"log"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Walls    []Wall
	Entities []Entity

	FOV int

	MainRaycaster *Raycaster

	DrawCalls []DrawCall

	Time int
}

var game = &Game{}

func (g *Game) Update() error {

	g.Time += 1

	for _, entity := range game.Entities {
		entity.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.DrawCalls = g.DrawCalls[:0]

	for _, entity := range game.Entities {
		entity.Draw(screen)
	}

	slices.SortFunc(g.DrawCalls, func(a, b DrawCall) int {
		return cmp.Compare(b.GetDepth(), a.GetDepth())
	})

	for _, drawCall := range g.DrawCalls {
		drawCall.Draw(screen)
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
	spawnEntity(&Book{
		StartPos: Vector2{
			320 / 2,
			240 / 2,
		},
	})
	//spawnEntity(&Chaser{})

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
