package editor

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/unix"
)


type cursor struct {
	row int
	col int
}

type Editor struct {
	lines []string
	cursor *cursor
	cols int
	rows int
	startingLinePos int
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
	e.startingLinePos = 0

	return e
}

func (e *Editor) GetCursor() *cursor {
	return e.cursor
}

func (e *Editor) printLines(pos int) {
    os.Stdout.Write([]byte("\033[2J")) //clear screen
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

func (e *Editor) RedrawScreen() {
    size, err := unix.IoctlGetWinsize(int(os.Stdin.Fd()), unix.TIOCGWINSZ)
	if err != nil {
        log.Fatal(err)
	}
	e.cols = int(size.Col)
	e.rows = int(size.Row)
	e.printLines(e.startingLinePos)
}

func (e *Editor) MoveCursorUp() {
	cursor := e.cursor
	fmt.Printf("row:%d start:%d\r", cursor.row, e.startingLinePos)
	if cursor.row == 0 && e.startingLinePos == 0 {
		return
	} else if cursor.row == 0 {
		e.startingLinePos -= 1
		e.RedrawScreen()
		return
	}
	cursor.row -= 1
	os.Stdout.Write([]byte("\033[1A"))
}

func (e *Editor) MoveCursorDown() {
	cursor := e.cursor 
	if (len(e.lines) - e.startingLinePos) < e.rows {
		return
	} else if cursor.row == (e.rows - 1) {
		e.startingLinePos += 1
		e.RedrawScreen()
		return
	} 
	cursor.row += 1
	os.Stdout.Write([]byte("\033[1B"))
}

func (c *cursor) MoveCursorLeft() {

}

func (c *cursor) MoveCursorRight() {

}