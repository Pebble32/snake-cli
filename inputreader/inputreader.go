package inputreader

import (
	"os"
)

type InputReader struct {
	buffer []byte
}

func New() *InputReader {
	return &InputReader{buffer: make([]byte, 1)}
}

func (ir* InputReader) Read(events chan byte) {
	for {
		for {
			readLen, _ := os.Stdin.Read(ir.buffer)
			if readLen > 0 {
				break
			}
		}
		if ir.buffer[0] == 27 {
			seq := make([]byte, 2)
			os.Stdin.Read(seq)
			if seq[0] == 91 {
				switch seq[1] {
					case 65:
						events <- 'w'
						continue
					case 66:
						events <- 's'
						continue
					case 67:
						events <- 'a'
						continue
					case 68:
						events <- 'd'
						continue
				}
			}
		}
		events <- ir.buffer[0]
	}
}
