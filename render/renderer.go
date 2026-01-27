package render

import (
	"fmt"
	"os"
	"snake/game"
)

type Renderer struct {}

const (
    clearScreen = "\x1b[2J"
    cursorHide = "\x1b[?25l"
    cursorShow = "\x1b[?25h"
    cursorHome = "\x1b[H"
)

func New() *Renderer{
	return &Renderer{}
}

func (r* Renderer) Render(g* game.Game) {
	strOut := clearScreen + cursorHide 
	for i, p := range g.Snake.Body {
		if len(g.Snake.Body) -1 == i {
            strOut += fmt.Sprintf("\x1b[%d;%dHO", p.Y+1, p.X+1)
			break
		}
        strOut += fmt.Sprintf("\x1b[%d;%dH@", p.Y+1, p.X+1)
	}

    strOut += fmt.Sprintf("\x1b[%d;%dH*", g.Food.Y+1, g.Food.X+1)
	os.Stdout.WriteString(strOut)
}

func (r* Renderer) Restore() {
	outString := cursorHide
	outString += clearScreen
	outString += cursorHome
	outString += cursorShow
	os.Stdout.WriteString(outString)
}
