package main

import "github.com/hajimehoshi/ebiten/v2"

type Raycaster struct {
	Position
	FOV int
}

func (r *Raycaster) update() {
	r.Position = game.Player.Position
}

func (r *Raycaster) draw(screen *ebiten.Image) {

}
