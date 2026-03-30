package main

import (
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Raycaster struct {
	BaseEntity
	Player *Player
}

var wallImage *ebiten.Image

func (r *Raycaster) Start() {
	game.MainRaycaster = r

	img, _, err := ebitenutil.NewImageFromFile("wall.png")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(img.Bounds())
	fmt.Println(img.Bounds().Dx())
	wallImage = img
}

type WallDrawCall struct {
	Slice   *ebiten.Image
	Options ebiten.DrawImageOptions
	Depth   int
}

func (w *WallDrawCall) Draw(screen *ebiten.Image) {
	screen.DrawImage(w.Slice, &w.Options)
}

func (w *WallDrawCall) GetDepth() int {
	return w.Depth
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

		brightness := max(float32(255-ray.Distance*2)/255, 0)

		wrappedX := wrap(int(ray.RelativePosition.X), 0, wallImage.Bounds().Dx()-1)

		slice := wallImage.SubImage(image.Rectangle{
			Min: image.Point{X: wrappedX, Y: 0},
			Max: image.Point{X: wrappedX + 1, Y: wallImage.Bounds().Dy()},
		}).(*ebiten.Image)
		options := ebiten.DrawImageOptions{}
		scaleY := float64(ySize) / float64(slice.Bounds().Dy())

		options.GeoM.Scale(1, scaleY)
		options.GeoM.Translate(float64(progress*360), float64((240/2)-ySize/2))
		options.Filter = ebiten.FilterNearest
		options.ColorScale.Scale(float32(brightness), float32(brightness), float32(brightness), 1)

		//screen.DrawImage(slice, &options)
		game.DrawCalls = append(game.DrawCalls, &WallDrawCall{slice, options, ray.Distance})
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

func wrap(n, min, max int) int {
	rangeSize := max - min + 1
	if rangeSize <= 0 {
		return min // or panic/error depending on your use case
	}

	// Normalize n into the range
	return ((n-min)%rangeSize+rangeSize)%rangeSize + min
}
