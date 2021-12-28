package burn

import (
	// "fmt"
)

type Position struct {
	Y int
	X int
}

func (pos *Position) InitPosition(y int, x int) Position{
	pos.Y = y
	pos.X = x

	return *pos
}