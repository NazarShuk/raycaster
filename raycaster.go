package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Raycaster struct {
	BaseEntity
	Player *Player
}

func (r *Raycaster) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%v", r.Player.Direction))

	for x := 0; x <= game.FOV; x++ {

		progress := float32(x) / float32(game.FOV)

		dir := Vector2{
			X: (progress - 0.5) * 2,
			Y: -1,
		}

		dir = rotate(dir, r.Player.Direction)

		ray := sendRay(r.Player.Position, dir, screen)

		ySize := float32(4000) / float32(ray.Distance)

		brightness := max(255-ray.Distance*2, 0)

		vector.FillRect(screen, progress*360, (240/2)-ySize/2, 1, ySize, color.RGBA{0, uint8(brightness), 0, 255}, false)

	}
}

type RayResult struct {
	Distance int
	Wall     *Wall

	RelativePosition Vector2
}

func sendRay(startPos Vector2, direction Vector2, screen *ebiten.Image) RayResult {

	rayPosition := Vector2{
		startPos.X,
		startPos.Y,
	}

	distance := 0

	var result OverlapResult
	var overlapWall *Wall

RayLoop:
	for {
		distance++

		rayPosition.X += direction.X
		rayPosition.Y += direction.Y

		//vector.FillRect(screen, rayPosition.X, rayPosition.Y, 2, 2, color.RGBA{255, 0, 0, 255}, true)

		for i := 0; i < len(game.Walls); i++ {

			result = checkOverlap(rayPosition, Vector2{1, 1}, game.Walls[i].Position, game.Walls[i].Size)
			if result.IsOverlapping {
				overlapWall = &game.Walls[i]
				break RayLoop
			}
		}

		if distance > 500 {

			break
		}

	}

	return RayResult{
		Distance:         distance,
		Wall:             overlapWall,
		RelativePosition: result.RelativePosition,
	}
}
