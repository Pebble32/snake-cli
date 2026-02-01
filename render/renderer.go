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

func (r *Renderer) RenderMenu(m *game.Menu, NCols, NRows int) {
	var strOut strings.Builder
	strOut.WriteString(clearScreen + cursorHide)

	logo := []string{
		`  _________ _______      _____   ____  __.___________`,
		` /   _____/ \      \    /  _  \ |    |/ _|\_   _____/`,
		` \_____  \  /   |   \  /  /_\  \|    |  <  |    __)_ `,
		` /        \/    |    \/    |    \    |  \  |        \`,
		`/_______  /\____|__  /\____|__  /____|__ \/_______  /`,
		`        \/         \/         \/        \/        \/`,
	}

	logoStartY := 2 
	for i, line := range logo {
		x := int(NCols/2) + 10		// columns
		y := logoStartY + i		// rows
		fmt.Fprintf(&strOut, "\x1b[%d;%dH%s", y, x, line)
	}

	for i, p := range m.Options {
		y := int(NCols/2) + i - 3
		x := int(NRows/2) - 3
		label := "  " + p
		if i == m.SelectedIndex {
			label = "> " + p
		}
		fmt.Fprintf(&strOut, "\x1b[%d;%dH%s", y, x, label)
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

	boxWidth := 34
	x := (NCols / 2) - (boxWidth / 2)
	y := (NRows / 2) - 3

	draw := func(rowOffset int, content string) {
		fmt.Fprintf(&strOut, "\x1b[%d;%dH%s", y+rowOffset, x, content)
	}

	draw(0, "+--------------------------------+")
	draw(1, "|           GAME OVER!           |")
	
	draw(2, fmt.Sprintf("|    YOUR SCORE WAS: %-5d       |", currentScore))
	
	draw(3, "+--------------------------------+")
	draw(4, "| Please input your name:        |")
	
	inputLine := fmt.Sprintf("> %s", name+"_")
	draw(5, fmt.Sprintf("| %-30s |", inputLine))
	
	draw(6, "+--------------------------------+")

	os.Stdout.WriteString(strOut.String())
}

func (r* Renderer) Restore() {
	outString := cursorHide
	outString += clearScreen
	outString += cursorHome
	outString += cursorShow
	os.Stdout.WriteString(outString)
}
