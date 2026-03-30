package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type BaseNPC struct {
	BaseEntity

	Position Vector2
	Sprite   *ebiten.Image
}

type BaseNPCDrawCall struct {
	DrawCall

	Sprite  *ebiten.Image
	screenX float32
	screenY float32
	scale   float32
}

func (d *BaseNPCDrawCall) Draw(screen *ebiten.Image) {
	if d.Sprite == nil {
		return
	}

	w, h := d.Sprite.Bounds().Dx(), d.Sprite.Bounds().Dy()

	op := &ebiten.DrawImageOptions{}

	scaleX := float64(d.scale) / float64(w)
	scaleY := float64(d.scale) / float64(h)
	op.GeoM.Scale(scaleX, scaleY)

	op.GeoM.Translate(float64(d.screenX)-float64(d.scale)/2, float64(d.screenY))

	screen.DrawImage(d.Sprite, op)
}

func (d *BaseNPCDrawCall) GetDepth() int {
	return int(4000 / d.scale)
}

func (c *BaseNPC) Draw(screen *ebiten.Image) {
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

	game.DrawCalls = append(game.DrawCalls, &BaseNPCDrawCall{
		Sprite:  c.Sprite,
		screenX: screenX,
		screenY: screenY,
		scale:   scale,
	})
}
