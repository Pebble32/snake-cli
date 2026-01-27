package game

type Snake struct {
	Body []Point
	Dir Direction
}

type Point struct {
	X int
	Y int
}

type Direction int
const (
	Up Direction = iota
	Down
	Left
	Right
)

func NewSnake() *Snake {
	head := Point{1,0}
	body := Point{0,0}

	return &Snake{
		Body: []Point{body, head},
		Dir: Right,
	}
}

func (s* Snake) Move() {
	head := s.Body[len(s.Body)-1]
	var newHead Point
	switch s.Dir {
	case Up:
		newHead = Point{head.X, head.Y - 1}
	case Down:
		newHead = Point{head.X, head.Y + 1}
	case Left:
		newHead = Point{head.X - 1, head.Y}
	case Right:
		newHead = Point{head.X + 1, head.Y}
	}
	s.Body = append(s.Body, newHead)
}

func (s* Snake) Pop() {
	s.Body = s.Body[1:]
}
