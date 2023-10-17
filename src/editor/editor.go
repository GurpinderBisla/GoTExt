package editor

import (
	"bufio"
	"log"
	"os"

	"golang.org/x/sys/unix"
)


type cursor struct {
	x int
	y int
}

type Editor struct {
	lines []string
	cursor *cursor
	cols int
	rows int
}

func NewEditor() *Editor {
	e := new(Editor)
	e.cursor = &cursor{0,0}

    size, err := unix.IoctlGetWinsize(int(os.Stdin.Fd()), unix.TIOCGWINSZ)
	if err != nil {
        log.Fatal(err)
	}

	e.cols = int(size.Col)
	e.rows = int(size.Row)

	return e
}

func (e *Editor) printLines(pos int) {
	for i := 0; pos < len(e.lines) && i < e.rows - 1; pos, i = pos + 1, i + 1 {
		os.Stdout.WriteString(e.lines[pos] + "\n\r")
	}
}

func (e *Editor) DrawUi() {
    os.Stdout.Write([]byte("\033[2J"))
    for i := 0; i < e.cols; i++ {
        if i != e.rows - 1 {
            os.Stdout.Write([]byte("~\n\r"))
        } else {
            os.Stdout.Write([]byte("~\033[0;0H"))
        }
    }
}

func (e *Editor) ReadFile(filepath string) {

	bufio.NewScanner(os.Stdin)
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		e.lines = append(e.lines, scanner.Text())
	}

	//print from the start of the file
	e.printLines(0)

	//return cursor to 0,0
	os.Stdout.WriteString("\033[0;0H")
}