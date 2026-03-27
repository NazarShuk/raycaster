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

	vector.FillRect(
		screen,
		screenX-scale/2,
		screenY,
		scale,
		scale,
		color.RGBA{uint8(scale), uint8(scale), uint8(scale), 255},
		false,
	)
}
