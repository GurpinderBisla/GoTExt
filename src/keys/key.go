package keys

import (
	"fmt"
	"goText/src/editor"
	"os"
)

const CTRLQ = 17
const BACKSPACE = 127
const ENTER = 13
const ARROW_ESC = 033

func ProcessKeyPress(c byte, ISRUNNING *bool, editor *editor.Editor) {
    switch c {
    case BACKSPACE: 
        output := []byte("\b \b")
        os.Stdout.Write(output)
    case CTRLQ:
        *ISRUNNING = false
    case ENTER:
        os.Stdout.Write([]byte("\n\r"))
    case ARROW_ESC:
            // if cBuff[0] == '[' {
            //     os.Stdout.Write([]byte("~\033[0;0H"))
            // }
    default:
        fmt.Printf("%c", c)
    }
}
