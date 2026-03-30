package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Book struct {
	BaseNPC

	StartPos Vector2
}

func (b *Book) Start() {
	img, _, err := ebitenutil.NewImageFromFile("book.png")
	if err != nil {
		panic(err)
	}

	b.Sprite = img
	b.Position = b.StartPos
}

func (b *Book) Update() {
	b.HeightOffset = float32(math.Sin(float64(game.Time)*0.1)) * 10
}
