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
    state := terminal.Startup()
    defer state.RestoreState()

    editor := editor.NewEditor()
    editor.DrawUi()
    editor.ReadFile("../test.txt")

    reader := bufio.NewReader(os.Stdin)
    for ISRUNNING := true; ISRUNNING;{
        c, err := reader.ReadByte()
        if err!= nil {
            log.Fatal(err)
        }
        keys.ProcessKeyPress(c, &ISRUNNING, editor)
    }
}
