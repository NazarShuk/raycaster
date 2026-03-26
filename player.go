package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	BaseEntity

	Position  Vector2
	Size      Vector2
	Direction float32
}

func (p *Player) Start() {
	raycaster := &Raycaster{Player: p}
	spawnEntity(raycaster)
}

func (p *Player) Draw(screen *ebiten.Image) {
	//vector.FillRect(screen, float32(p.Position.X), float32(p.Position.Y), float32(p.Size.X), float32(p.Size.Y), color.RGBA{255, 255, 255, 255}, true)
}

func (p *Player) Update() {

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
