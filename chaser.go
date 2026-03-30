package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Chaser struct {
	BaseEntity

	Position Vector2
}

type ChaserDrawCall struct {
	DrawCall,
	screenX float32
	screenY float32
	scale   float32
}

func (d *ChaserDrawCall) Draw(screen *ebiten.Image) {
	vector.FillRect(
		screen,
		d.screenX-d.scale/2,
		d.screenY,
		d.scale,
		d.scale,
		color.RGBA{uint8(d.scale), uint8(d.scale), uint8(d.scale), 255},
		false,
	)
}

func (d *ChaserDrawCall) GetDepth() int {
	return int(4000 / d.scale)
}

func (c *Chaser) Update() {
	c.Position = c.Position.lerp(game.MainRaycaster.Player.Position, 0.01)
}

func (c *Chaser) Draw(screen *ebiten.Image) {
	diffX := game.MainRaycaster.Player.Position.X - c.Position.X
	diffY := game.MainRaycaster.Player.Position.Y - c.Position.Y

	screenPosition := Vector2{diffX, diffY}
	screenPosition = rotate(screenPosition, -game.MainRaycaster.Player.Direction)

	scale := 4000 / screenPosition.Y

	if scale < 0 {
		return
	}

	screenX := 320/2 - (screenPosition.X/screenPosition.Y)*240
	screenY := float32(480)/4 - scale/2

	game.DrawCalls = append(game.DrawCalls, &ChaserDrawCall{
		screenX: screenX,
		screenY: screenY,
		scale:   scale,
	})
}
