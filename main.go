package main

import (
	"math"
	"snake/game"
	"snake/inputreader"
	"snake/render"
	"snake/terminal"
	"time"
)

func atExit(t* terminal.Terminal, r* render.Renderer, ticker* time.Ticker) {
	t.Restore()
	r.Restore()
	ticker.Stop()
}

func getUserName(events chan byte, r *render.Renderer, score, NCol, NRow int) string {
	name := ""
	r.RenderInputNameScreen(name, score, NCol, NRow)
	for {
		char := <- events

		switch char {
		case '\r', '\n':
			return name
		case 127, 8:
			if len(name) > 0{
				name = name[:len(name) - 1] 
			}
		default:
			if len(name) < 23 {
				name += string(char)
			}
		}
		r.RenderInputNameScreen(name, score, NCol, NRow)
	}
}

func main(){
	t, err := terminal.New()
	if err != nil {
		panic(err)
	}

	ir := inputreader.New()
	g := game.New(t.NRows, t.NCols)
	m := game.NewMenu()
	r := render.New()
	sc := game.NewScore()

	tickRate := time.Second / 10
	ticker := time.NewTicker(tickRate)
	events := make(chan byte)
	go ir.Read(events)
	var input byte

	defer atExit(t, r, ticker)

	r.RenderMenu(m, t.NRows, t.NCols)
	for {
		input = <- events	
		result := m.Update(input)
		r.RenderMenu(m, t.NRows, t.NCols)

		switch result {
		case game.StartGame:
			for {
				select {
				case <- ticker.C:
					score1 := len(g.Snake.Body)
					r.RenderSnake(g)
					g.Update(input)
					score2 := len(g.Snake.Body)
					if score2 > score1 && score2 < 11 {
						ticker.Stop()
						multiplier := 1.0 + (math.Log10(float64(score2) + 1.0))
						tickRate = tickRate / time.Duration(multiplier)
						ticker = time.NewTicker(tickRate)
					}
					if g.GameOver() {
						name := getUserName(events, r, len(g.Snake.Body), t.NCols, t.NRows)
						sc.SaveScore(name, len(g.Snake.Body))
						return
					}
				case input = <-events:
					if input == 'q' {
						return
					}
				}
			}
		case game.HighScore:
			r.RenderHighScore(sc, t.NCols, t.NRows)
		case game.Exit:
			return
		}
	}
}
