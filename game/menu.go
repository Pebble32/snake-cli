package game

type Menu struct {
	Options       []string
	SelectedIndex int
}

const (
	StartGame = iota
	HighScore
	Exit
	numOptions
)

func NewMenu() *Menu {
	return &Menu{Options: []string{"Start Game", "High Score", "Exit"}, SelectedIndex: StartGame}
}

func (m *Menu) Update(key byte) int {
	switch key {
	case 'w':
		if m.SelectedIndex > StartGame {
			m.SelectedIndex--
		} else {
			m.SelectedIndex = Exit
		}
	case 's':
		if m.SelectedIndex < numOptions {
			m.SelectedIndex++
		} else {
			m.SelectedIndex = StartGame
		}
	case '\r', ' ', '\n':
		return m.SelectedIndex
	}

	return -1
}
