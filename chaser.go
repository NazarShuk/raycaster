package main

import "github.com/hajimehoshi/ebiten/v2/ebitenutil"

type Chaser struct {
	BaseNPC
}

func (c *Chaser) Start() {
	img, _, err := ebitenutil.NewImageFromFile("chaser.png")
	if err != nil {
		panic(err)
	}

	c.Sprite = img
}

func (c *Chaser) Update() {
	c.Position = c.Position.lerp(game.MainRaycaster.Player.Position, 0.01)
}
