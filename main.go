package main

import (
	"snake/terminal"
	"snake/inputreader"
)

func main(){
	t, err := terminal.New()
	if err != nil {
		panic(err)
	}

	defer t.Restore()

	ir := inputreader.New()

	for {
		key := ir.Read()

		if key == 'q' {
			break
		}
	}
}
