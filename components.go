package main

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
