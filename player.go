package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	Position  Vector2
	Size      Vector2
	Direction float32
}

func (p *Player) draw(screen *ebiten.Image) {
	vector.FillRect(screen, float32(p.Position.X), float32(p.Position.Y), float32(p.Size.X), float32(p.Size.Y), color.RGBA{255, 255, 255, 255}, true)
}

func (p *Player) update() {

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.Direction -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.Direction += 1
	}

	movementChange := Vector2{}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		movementChange.Y = -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		movementChange.Y = 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		movementChange.X = 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		movementChange.X = -1
	}

	movementChange = rotate(movementChange, p.Direction)

	p.Position.add(movementChange)
}
