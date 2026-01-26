package terminal

import (
	"os"
	"syscall"
	"unsafe"
)

type Terminal struct {
	fd uintptr
	original Termios
	modified *Termios
	NCols int
	NRows int
}

func New() (*Terminal, error) {
	t := &Terminal{}

	t.fd = os.Stdout.Fd()
	termios, err := getTermios(t.fd)
	if err != nil {
		return nil, err
	}

	t.original = *termios
	t.modified = termios

	t.enableRawMode()
	err = t.GetWindowSize()
	if err != nil {
		return nil, err
	}
	
	return t, nil
}

func (t* Terminal) Restore() error {
	return setTermios(t.fd, &t.original)
}

func (t* Terminal) GetWindowSize() error {
	ws := struct {
		rows uint16 
		columns uint16
	}{}

	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL, // I want to talk to a device 
		t.fd, // This Specific device
		syscall.TIOCGWINSZ, // Specifically Terminal size  
		uintptr(unsafe.Pointer(&ws)), // Write information here
	) 
	if errno != 0 {
		return errno
	}

	t.NCols = int(ws.columns)
	t.NRows = int(ws.rows)

	return nil
}

func (t* Terminal) enableRawMode() error{
    t.modified.Lflag &^= syscall.ECHO | syscall.ICANON | syscall.ISIG | syscall.IEXTEN
    t.modified.Iflag &^= syscall.BRKINT | syscall.ICRNL | syscall.INPCK | syscall.ISTRIP | syscall.IXON
    t.modified.Cflag |= syscall.CS8
    t.modified.Oflag &^= syscall.OPOST
    t.modified.Cc[syscall.VMIN+1] = 0
    t.modified.Cc[syscall.VTIME+1] = 1
	return setTermios(t.fd, t.modified)
} 
