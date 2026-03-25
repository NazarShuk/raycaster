package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Position struct {
	X int
	Y int
}

type Size struct {
	X int
	Y int
}

type Player struct {
	Position
	Size
}

func (p *Player) draw(screen *ebiten.Image) {
	vector.FillRect(screen, float32(p.Position.X), float32(p.Position.Y), float32(p.Size.X), float32(p.Size.Y), color.RGBA{255, 255, 255, 255}, true)
}

func (p *Player) update() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Position.Y -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Position.Y += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Position.X += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Position.X -= 1
	}
}
