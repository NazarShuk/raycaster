package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vector2 struct {
	X float32
	Y float32
}

type OverlapResult struct {
	IsOverlapping    bool
	RelativePosition Vector2
}

func isOverlapping(sourcePos Vector2, sourceSize Vector2, comparePos Vector2, compareSize Vector2) bool {
	return sourcePos.X < comparePos.X+compareSize.X &&
		sourcePos.X+sourceSize.X > comparePos.X &&
		sourcePos.Y < comparePos.Y+compareSize.Y &&
		sourcePos.Y+sourceSize.Y > comparePos.Y
}

func checkOverlap(sourcePos Vector2, sourceSize Vector2, comparePos Vector2, compareSize Vector2) OverlapResult {
	isOverlapping := sourcePos.X < comparePos.X+compareSize.X &&
		sourcePos.X+sourceSize.X > comparePos.X &&
		sourcePos.Y < comparePos.Y+compareSize.Y &&
		sourcePos.Y+sourceSize.Y > comparePos.Y

	var relativePos Vector2

	if isOverlapping {
		relativePos.X = sourcePos.X - comparePos.X
		relativePos.Y = sourcePos.Y - comparePos.Y
	}

	return OverlapResult{
		IsOverlapping:    isOverlapping,
		RelativePosition: relativePos,
	}
}

func rotate(v Vector2, degrees float32) Vector2 {
	angle := degrees * math.Pi / 180

	cos := float32(math.Cos(float64(angle)))
	sin := float32(math.Sin(float64(angle)))

	return Vector2{
		X: v.X*cos - v.Y*sin,
		Y: v.X*sin + v.Y*cos,
	}
}

func (v *Vector2) add(anotherVector Vector2) {
	v.X += anotherVector.X
	v.Y += anotherVector.Y
}

func (v *Vector2) lerp(b Vector2, t float32) Vector2 {
	return Vector2{
		X: lerp(v.X, b.X, t),
		Y: lerp(v.Y, b.Y, t),
	}
}

func lerp(a float32, b float32, t float32) float32 {
	return a + (b-a)*t
}

type Entity interface {
	Start()
	Draw(screen *ebiten.Image)
	Update()
}

type BaseEntity struct{}

func (b *BaseEntity) Start()                    {}
func (b *BaseEntity) Update()                   {}
func (b *BaseEntity) Draw(screen *ebiten.Image) {}

type DrawCall interface {
	Draw(screen *ebiten.Image)
	GetDepth() int
}
