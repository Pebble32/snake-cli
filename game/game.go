package game

import "math/rand"

type Game struct {
	NRows int
	NCols int
	Snake *Snake
	Food Point
}

func New(nRows, nCols int) *Game{
	g := &Game{NRows: nRows, NCols: nCols}
	g.Snake = NewSnake()
	g.spawnFood()

	return g
}

func (g* Game) Update(key byte) {
	switch key{
	case 'w':
		g.Snake.Dir = Up
	case 's':
		g.Snake.Dir = Down
	case 'a':
		g.Snake.Dir = Left
	case 'd':
		g.Snake.Dir = Right
	}
	g.Snake.Move()
	head := g.Snake.Body[len(g.Snake.Body) - 1]
	if head == g.Food {
		g.spawnFood()
	} else {
		g.Snake.Pop()
	}
}

func (g* Game) spawnFood() {
	x := rand.Intn(g.NCols)
	y := rand.Intn(g.NRows)

	g.Food =  Point{x, y}

	for _, p := range g.Snake.Body {
		if p == g.Food {
			g.spawnFood()
		}
	}
}

func (g* Game) GameOver() bool{
	head := g.Snake.Body[len(g.Snake.Body) - 1]
	if head.X < 0 || head.X > g.NCols || head.Y < 0 || head.Y > g.NRows {
		return true
	}

	for i, p := range g.Snake.Body {
		if i == len(g.Snake.Body)-1 {
			continue
		}
		if p == head {
			return true
		}
	}

	return false
}
