package main

import (
	"bufio"
	"os"
	"os/exec"

	menu "github.com/MisaelCodes/code_learning_scheduler/components"
	"golang.org/x/term"
)



func main() {
	// start with the menu options
	menuOptions := []string{
		"Show List",
		"Create List",
	}

	// Put the terminal in raw input mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

    // Create init and render the menu
	prompt := "Move with 'j' and 'k', press q to quit:"
	m := menu.NewMenu(bufio.NewReader(os.Stdin), prompt, menuOptions)
	m.Render()
    command := "cd ~/Documents/learning/golang/neetcode_algs && nvim main.go"
    cmd := exec.Command("gnome-terminal","--", "bash", "-c", command)
    if err := cmd.Run(); err != nil{
        panic(err)
    }

}
