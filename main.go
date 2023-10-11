package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pkg/term/termios"
	"golang.org/x/sys/unix"
)

const CTRLQ = 17

type terminalState struct {
    oldState *unix.Termios
    newState *unix.Termios
}

func newTerminalState() *terminalState {
    p := new(terminalState)
    p.oldState = &unix.Termios{}
    p.newState = &unix.Termios{}
    return p
}


// TODO set flags for TCSADRAIN and TCSAFLUSH
func(t *terminalState) RestoreState() {
    err := termios.Tcsetattr(os.Stdin.Fd(), termios.TCSAFLUSH, t.oldState)
    if err != nil {
        log.Fatal(err)
    }
}

func(t *terminalState) MakeRaw() {
    err := termios.Tcgetattr(os.Stdin.Fd(), t.oldState)
    if err != nil {
        log.Fatal(err)
    }
    //creating copy oldState, doing it in a roundabout way to evoid warning
    tmpNewState := *t.oldState
    t.newState = &tmpNewState

    // direct implemtation of the man page for rawterminals
    t.newState.Iflag &^= (unix.IGNBRK | unix.BRKINT | unix.PARMRK | unix.ISTRIP | unix.INLCR | unix.IGNCR | unix.ICRNL | unix.IXON)
    t.newState.Oflag &^= (unix.OPOST)
    t.newState.Lflag &^= (unix.ECHO | unix.ECHONL | unix.ICANON | unix.ISIG | unix.IEXTEN)
    t.newState.Cflag &^= (unix.CSIZE | unix.PARENB)
    t.newState.Cflag |= unix.CS8
    t.newState.Cc[unix.VMIN] = 1
    t.newState.Cc[unix.VTIME] = 0

    err = termios.Tcsetattr(os.Stdin.Fd(), termios.TCSAFLUSH, t.newState)
    if err != nil {
        log.Fatal(err)
    }
}

func printToTerminalTest() {
    test := "Hello World\n\rBonjour World \n\r"
    fmt.Print(test)
}

func main() {
    state := newTerminalState() 
    state.MakeRaw()
    defer state.RestoreState()

    var c byte
    reader := bufio.NewReader(os.Stdin)
    printToTerminalTest()
    for c != CTRLQ {
        tmpByte, err := reader.ReadByte()
        if err!= nil {
            log.Fatal(err)
        }
        c = tmpByte
        fmt.Printf("%c:%d\n\r", c, c)
    }
    fmt.Printf("Done\n\r")
}
