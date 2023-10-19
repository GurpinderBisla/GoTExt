package keys

import (
	"fmt"
	"goText/src/editor"
	"os"
)

const CTRLQ = 17
const BACKSPACE = 127
const ENTER = 13
const ARROW_UP = 'A'
const ARROW_DOWN = 'B'
const ARROW_RIGHT = 'C'
const ARROW_LEFT = 'D'

func ProcessKeyPress(c byte, ISRUNNING *bool, editor *editor.Editor) {
    // reader := bufio.NewReader(os.Stdin)
    switch c {
    case BACKSPACE: 
        output := []byte("\b \b")
        os.Stdout.Write(output)
    case CTRLQ:
        *ISRUNNING = false
    case ENTER:
        os.Stdout.Write([]byte("\n\r"))
    case ARROW_UP:
        editor.MoveCursorUp()
    case ARROW_DOWN:
        editor.MoveCursorDown()
    case ARROW_LEFT:
        os.Stdout.Write([]byte("\033[1D"))
    case ARROW_RIGHT:
        os.Stdout.Write([]byte("\033[1C"))
    default:
        fmt.Printf("%c", c)
    }
}
