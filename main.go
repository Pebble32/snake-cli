package main

import (
	"os"
	"snake/terminal"
)

func main(){
	t, err := terminal.New()
	if err != nil {
		panic(err)
	}

	defer t.Restore()

	buffer := make([]byte, 1)
	for {
		os.Stdin.Read(buffer)

		if buffer[0] == 'q' {
			break
		}
	}
}
