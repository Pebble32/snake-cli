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

func (r* Renderer) RenderSnake(g* game.Game) {
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

func (r *Renderer) RenderHighScore(sc *game.Score, NCols, NRows int) {
	allScores := sc.LoadSortedScores()
	var strOut strings.Builder
	strOut.WriteString(clearScreen + cursorHide)

	const (
		boxWidth   = 46
		innerWidth = boxWidth - 2 
		nameWidth  = 20
	)

	startX := (NCols / 2) - (boxWidth / 2)
	startY := (NRows / 2) - 6

	lineBorder := "+--------------------------------------------+"

	fmt.Fprintf(&strOut, "\x1b[%d;%dH%s", startY, startX, lineBorder)

	title := "SNAKE HIGH SCORES"
	leftPad := (innerWidth - len(title)) / 2
	rightPad := innerWidth - len(title) - leftPad
	fmt.Fprintf(
		&strOut,
		"\x1b[%d;%dH|%s%s%s|",
		startY+1,
		startX,
		strings.Repeat(" ", leftPad),
		title,
		strings.Repeat(" ", rightPad),
	)

	fmt.Fprintf(&strOut, "\x1b[%d;%dH%s", startY+2, startX, lineBorder)

	// Score top 10 for now
	for i := range 10 {
		currentY := startY + 3 + i

		var line string
		if i < len(allScores) {
			p := allScores[i]

			displayName := p.Name
			if len(displayName) > nameWidth {
				displayName = displayName[:nameWidth-3] + "..."
			}

			// 43 chars total -> last char padded to 44 for breathing space
			line = fmt.Sprintf(
				"%2d. %-20s %5dpts %s",
				i+1,
				displayName,
				p.Score,
				p.Date.Format("2006-01-02"),
			)
		}

		if len(line) > innerWidth {
			line = line[:innerWidth]
		} else {
			line += strings.Repeat(" ", innerWidth-len(line))
		}

		fmt.Fprintf(&strOut, "\x1b[%d;%dH|%s|", currentY, startX, line)
	}

	lastLine := startY + 13
	fmt.Fprintf(&strOut, "\x1b[%d;%dH%s", lastLine, startX, lineBorder)

	instruction := "Press any key to return to menu..."
	instrX := (NCols / 2) - (len(instruction) / 2)
	fmt.Fprintf(&strOut, "\x1b[%d;%dH%s", lastLine+2, instrX, instruction)

	os.Stdout.WriteString(strOut.String())
}

// Helper function because Go's built-in math.Max uses float64
// I need int to know the offset
func max(a, b int) int {
	if a > b { return a }
	return b
}

func (r *Renderer) RenderInputNameScreen(name string, currentScore, NCols, NRows int) {
	var strOut strings.Builder
	strOut.WriteString(clearScreen + cursorHide)

	boxWidth := 33
	boxHeight := 6
	x := (NCols / 2) - (boxWidth / 2)
	y := (NRows / 2) - (boxHeight / 2)

	fmt.Fprintf(&strOut, "\x1b[%d;%dH+--------------------------------+", y, x)
	fmt.Fprintf(&strOut, "\x1b[%d;%dH|        GAME OVER!              |", y+1, x)
	fmt.Fprintf(&strOut, "\x1b[%d;%dH|    YOUR SCORE WAS: %-5d       |", y+2, x, currentScore)
	fmt.Fprintf(&strOut, "\x1b[%d;%dH+--------------------------------+", y+3, x)
	fmt.Fprintf(&strOut, "\x1b[%d;%dH| Please input your name:        |", y+4, x)
	fmt.Fprintf(&strOut, "\x1b[%d;%dH| > %-26s |", y+5, x, name+"_")
	fmt.Fprintf(&strOut, "\x1b[%d;%dH+--------------------------------+", y+6, x)
	os.Stdout.WriteString(strOut.String())
}

func (r* Renderer) Restore() {
	outString := cursorHide
	outString += clearScreen
	outString += cursorHome
	outString += cursorShow
	os.Stdout.WriteString(outString)
}
