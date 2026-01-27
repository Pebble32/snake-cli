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

func (ir* InputReader) Read() byte{
	os.Stdin.Read(ir.buffer)

	if ir.buffer[0] == 27 {
		seq := make([]byte, 2)
		os.Stdin.Read(seq)
		if seq[0] == 91 {
			switch seq[1] {
				case 65:
					return 'w'
				case 66:
					return 's'
				case 67:
					return 'a'
				case 68:
					return 'd'
			}
		}
	}
	return ir.buffer[0]
}
