package render

import (
	"strings"
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
	var strOut strings.Builder; strOut.WriteString(clearScreen + cursorHide) 
	for i, p := range g.Snake.Body {
		if len(g.Snake.Body) -1 == i {
            fmt.Fprintf(&strOut, "\x1b[%d;%dHO", p.Y+1, p.X+1)
			break
		}
        fmt.Fprintf(&strOut, "\x1b[%d;%dH@", p.Y+1, p.X+1)
	}

    fmt.Fprintf(&strOut, "\x1b[%d;%dH*", g.Food.Y+1, g.Food.X+1)
	os.Stdout.WriteString(strOut.String())
}

func (r* Renderer) RenderMenu(m* game.Menu, NCols, NRows int) {
	var strOut strings.Builder; strOut.WriteString(clearScreen + cursorHide)
	for i, p := range m.Options {
		x := int(NCols / 2) + i - 3
		y := int(NRows / 2) - 3
		label := p
		if i == m.SelectedIndex {
			label = "> " + p + "\n"
		} else {
			label = "  " + p + "\n"
		}
		// We add this specific line to the string in this specific position
		fmt.Fprintf(&strOut, "\x1b[%d;%dH%s", x, y, label)
	}
	os.Stdout.WriteString(strOut.String())
}

func (r* Renderer) RenderHighScore() {
	
}

func (r* Renderer) Restore() {
	outString := cursorHide
	outString += clearScreen
	outString += cursorHome
	outString += cursorShow
	os.Stdout.WriteString(outString)
}
