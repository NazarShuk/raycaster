package main

import "math"

type Vector2 struct {
	X float32
	Y float32
}

func isOverlapping(sourcePos Vector2, sourceSize Vector2, comparePos Vector2, compareSize Vector2) bool {
	return sourcePos.X < comparePos.X+compareSize.X &&
		sourcePos.X+sourceSize.X > comparePos.X &&
		sourcePos.Y < comparePos.Y+compareSize.Y &&
		sourcePos.Y+sourceSize.Y > comparePos.Y
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
