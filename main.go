package main

import (
	"snake/game"
	"snake/inputreader"
	"snake/render"
	"snake/terminal"
	"time"
	"fmt"
)

func atExit(t* terminal.Terminal, r* render.Renderer, ticker* time.Ticker) {
	t.Restore()
	r.Restore()
	ticker.Stop()
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
						return
					}
				case input = <-events:
					if input == 'q' {
						return
					}
				}
			}
		case game.HighScore:
			fmt.Println("This is high score")
		case game.Exit:
			return
		}

	}
}
