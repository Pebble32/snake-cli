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

func main(){
	t, err := terminal.New()
	if err != nil {
		panic(err)
	}




	ir := inputreader.New()

	g := game.New(t.NRows, t.NCols)

	r := render.New()

	tickRate := time.Second / 10
	ticker := time.NewTicker(tickRate)
	events := make(chan byte)
	go ir.Read(events)
	var input byte

	defer atExit(t, r, ticker)
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
}
