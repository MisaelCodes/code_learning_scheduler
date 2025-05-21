package main

import (
	"bufio"
	"fmt"
	"os"

	menu "github.com/MisaelCodes/code_learning_scheduler/components"
	"github.com/MisaelCodes/code_learning_scheduler/logic"
	"golang.org/x/term"
)



func main() {
	// start with the menu options
	menuOptions := []string{
		"Show List",
		"Create List",
	}

    tl := logic.TaskList{}
    r, err := tl.Fields()
    if err != nil{
        panic(err)
    }
    fmt.Println(r)
    
	// Put the terminal in raw input mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

    // Create init and render the menu
	prompt := "Move with 'j' and 'k', press q to quit:"
	m := menu.NewMenu(bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)), prompt, menuOptions)
	m.Render()
    // command := "cd ~/Documents/learning/golang/neetcode_algs && nvim main.go"
    // cmd := exec.Command("gnome-terminal","--", "bash", "-c", command)
    // if err := cmd.Run(); err != nil{
    //     panic(err)
    // }

}
