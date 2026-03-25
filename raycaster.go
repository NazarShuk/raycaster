package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Raycaster struct {
	Position Vector2
	FOV      int
}

func (r *Raycaster) update() {
	r.Position = game.Player.Position

}

func (r *Raycaster) draw(screen *ebiten.Image) {
	//wg := sync.WaitGroup{}

	for x := 0; x < game.FOV; x++ {

		//wg.Go(func() {
		progress := float32(x) / float32(game.FOV)

		distance := sendRay(r.Position, Vector2{
			X: (progress - 0.5) * 2,
			Y: -1,
		}, screen)

		ySize := float32(4000) / float32(distance)

		brightness := max(255-distance*2, 0)

		vector.FillRect(screen, progress*360, (240/2)-ySize/2, 2, ySize, color.RGBA{0, uint8(brightness), 0, 255}, false)
		//})

	}
	//wg.Wait()
}

func sendRay(startPos Vector2, direction Vector2, screen *ebiten.Image) int {

	rayPosition := Vector2{
		startPos.X,
		startPos.Y,
	}

	distance := 0

RayLoop:
	for {
		distance++

		rayPosition.X += direction.X
		rayPosition.Y += direction.Y

		//vector.FillRect(screen, rayPosition.X, rayPosition.Y, 2, 2, color.RGBA{255, 0, 0, 255}, true)

		for i := 0; i < len(game.Walls); i++ {
			if (isOverlapping(rayPosition, Vector2{1, 1}, game.Walls[i].Position, game.Walls[i].Size)) {
				break RayLoop
			}
		}

		if distance > 500 {

			break
		}

	}

	return distance
}
