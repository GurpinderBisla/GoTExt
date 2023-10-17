package main

import (
	"bufio"
	"goText/src/editor"
	"goText/src/keys"
	"goText/src/terminal"
	"log"
	"os"
)

func main() {
    setup()
}


//TODO move everything below to their own files
func setup() {
    state := terminal.Startup()
    defer state.RestoreState()

    editor := editor.NewEditor()
    editor.DrawUi()
    editor.ReadFile("../test.txt")
    eventLoop(editor)
}

func eventLoop(editor *editor.Editor) {
    reader := bufio.NewReader(os.Stdin)
    for ISRUNNING := true; ISRUNNING;{
        c, err := reader.ReadByte()
        fatalError(err)

        //Test for arrow keys or single ESC code (033 = ESC, arrows = \033[<A/B/C/D>, 
        // where ABCD are the arrow directions)
        if c == 033 {
            arrowTest, err := reader.ReadByte()
            fatalError(err)

            if arrowTest == '[' {
                arrowDirection, err := reader.ReadByte()
                fatalError(err)
                keys.ProcessKeyPress(arrowDirection, &ISRUNNING, editor)
                continue
            }
        }
        keys.ProcessKeyPress(c, &ISRUNNING, editor)
    }
}

func fatalError(e error) {
    if e !=nil {
        log.Fatal(e)
    }
}