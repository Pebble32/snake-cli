package main

import (
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

func getUserName(events chan byte, r *render.Renderer) string {
	name := ""
	r.RenderInputNameScreen(name)
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
		r.RenderInputNameScreen(name)
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
					r.Render(g)
					g.Update(input)
					if g.GameOver() {
						name := getUserName(events, r)
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
			r.RenderHighScore(sc)
		case game.Exit:
			return
		}
	}
}
