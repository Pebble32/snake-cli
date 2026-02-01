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

func (r *Renderer) RenderHighScore(sc *game.Score, NCols, NRows int) {
	allScores := sc.LoadSortedScores()
	var strOut strings.Builder
	strOut.WriteString(clearScreen + cursorHide)

	// 1. Define Box Width and find Center
	boxWidth := 40
	startX := (NCols / 2) - (boxWidth / 2)
	startY := (NRows / 2) - 5 // Offset upward so the list starts near the middle

	// 2. Draw Header
	fmt.Fprintf(&strOut, "\x1b[%d;%dH+--------------------------------------------+", startY, startX)
	fmt.Fprintf(&strOut, "\x1b[%d;%dH|          SNAKE HIGH SCORES           |", startY+1, startX)
	fmt.Fprintf(&strOut, "\x1b[%d;%dH+--------------------------------------------+", startY+2, startX)

	// 3. Draw Scores (Limit to Top 10)
	for i, p := range allScores {
		if i >= 10 { break } // Don't overflow the screen
		
		// Create a formatted string for the player info
		// %-3d = Rank (3 spaces)
		// %-12s = Name (12 spaces, left aligned)
		// %6d = Score (6 spaces, right aligned)
		scoreLine := fmt.Sprintf("%2d. %-12s %6d pts  %s", 
			i+1, p.Name, p.Score, p.Date.Format("2006-01-02"))

		fmt.Fprintf(&strOut, "\x1b[%d;%dH| %-36s |", startY+3+i, startX, scoreLine)
	}

	// 4. Fill empty lines if there are fewer than 5 scores
	if len(allScores) < 5 {
		for i := len(allScores); i < 5; i++ {
			fmt.Fprintf(&strOut, "\x1b[%d;%dH|                                      |", startY+3+i, startX)
		}
	}

	// 5. Draw Footer
	lastLine := startY + 3 + max(len(allScores), 5)
	fmt.Fprintf(&strOut, "\x1b[%d;%dH+--------------------------------------------+", lastLine, startX)
	fmt.Fprintf(&strOut, "\x1b[%d;%dH  Press any key to return to menu...", lastLine+2, startX+2)

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
