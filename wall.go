package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Wall struct {
	BaseEntity
	Position Vector2
	Size     Vector2
}

func (w *Wall) Draw(screen *ebiten.Image) {
	//vector.FillRect(screen, float32(w.Position.X), float32(w.Position.Y), float32(w.Size.X), float32(w.Size.Y), color.RGBA{128, 128, 128, 255}, true)
}
