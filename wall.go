package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Wall struct {
	Position
	Size
}

func (w *Wall) draw(screen *ebiten.Image) {
	vector.FillRect(screen, float32(w.Position.X), float32(w.Position.Y), float32(w.Size.X), float32(w.Size.Y), color.RGBA{128, 128, 128, 255}, true)
}
