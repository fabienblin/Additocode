package burn

import (
	"math/rand"
	"os"
)

type Field struct {
	Y int
	X int
	Grid [][] rune
}

func (f *Field) InitField(y int, x int) Field{
	f.Y = y
	f.X = x

	var grid [][] rune = make([][]rune, y)
	for i := 0; i < y; i++ {
		grid[i] = make([]rune, x)
	}
	f.Grid = grid

	return *f
}

func (f *Field) setSymbolOnPosition(pos Position, r rune) {
	if (!f.isOutOfBoundries(pos)){
		f.Grid[pos.Y][pos.X] = r
	}
}

func (f *Field) getSymbolOnPosition(pos Position) rune {
	if (!f.isOutOfBoundries(pos)) {
		return f.Grid[pos.Y][pos.X]
	}
	return 0
}

func (f *Field) isOutOfBoundries(pos Position) bool {
	return (pos.X < 0 ||
			pos.Y < 0 ||
			pos.X >= f.X ||
			pos.Y >= f.Y)
}

func (f *Field) InitWeed() {
	for i, _ := range f.Grid {
		for j, _ := range f.Grid[i] {
			f.setSymbolOnPosition(new(Position).InitPosition(i, j), WEED)
		}
	}
}

func (f *Field) InitFires(fires [] Position) {
	for _, fire := range fires {
		f.setSymbolOnPosition(fire, FIRE)
	}
}

func (f *Field) PropagateFire(pos Position) {
	if (f.getSymbolOnPosition(pos) == WEED && rand.Float64() >= Env.PropagationRate){
		f.setSymbolOnPosition(pos, FIRE)
	}
}

func (f *Field) getFires() []Position{
	var fires []Position

	for i, _ := range f.Grid {
		for j, _ := range f.Grid[i] {
			if (f.Grid[i][j] == FIRE) {
				fires = append(fires, new(Position).InitPosition(i, j))
			}
		}
	}

	return fires
}

func (f *Field) SimulateFires() {
	Env.Fires = f.getFires()

	if (Env.Fires == nil){
		os.Exit(0)
	}

	for _, fire := range Env.Fires {
		f.setSymbolOnPosition(fire, ASHES)

		up := new(Position).InitPosition(fire.Y - 1, fire.X)
		down := new(Position).InitPosition(fire.Y + 1, fire.X)
		left := new(Position).InitPosition(fire.Y, fire.X - 1)
		right := new(Position).InitPosition(fire.Y, fire.X + 1)

		f.PropagateFire(up)
		f.PropagateFire(down)
		f.PropagateFire(left)
		f.PropagateFire(right)
	}

	Env.Fires = nil
}