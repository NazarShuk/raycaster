package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Raycaster struct {
	Position Vector2
	FOV      int
}

func (r *Raycaster) update() {
	if ebiten.IsKeyPressed(ebiten.Key0) {
		r.Position = game.Player.Position

	}
}

func (r *Raycaster) draw(screen *ebiten.Image) {
	rayPosition := Vector2{
		X: r.Position.X + 8,
		Y: r.Position.Y + 8,
	}
	rayDirection := Vector2{
		X: 0,
		Y: -1,
	}

	distance := 0

RayLoop:
	for {
		distance++

		rayPosition.X += rayDirection.X
		rayPosition.Y += rayDirection.Y

		vector.FillRect(screen, rayPosition.X, rayPosition.Y, 2, 2, color.RGBA{255, 0, 0, 255}, true)

		for i := 0; i < len(game.Walls); i++ {
			if (isOverlapping(rayPosition, Vector2{1, 1}, game.Walls[i].Position, game.Walls[i].Size)) {
				break RayLoop
			}
		}

		if distance > 500 {

			ebitenutil.DebugPrint(screen, "Couldn't collide")
			break
		}

	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%v", distance))

}
