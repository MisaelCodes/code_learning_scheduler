package components

import "fmt"

const (
	ESC           = "\033"
	HIDE_CURSOR   = "[?25l"
	SHOW_CURSOR   = "[?25h"
	LOCATE_CURSOR = "[6n"
	NEXT_LINE     = "[1E"
	RESET         = "[0m"
)

// Formatters

// moves cursor to the line in position l
func MoveCursor(l int) {
	fmt.Printf("%s[%d;0H", ESC, l)
}

func SetForegroundColor(c int) error {
    if c > 255 || c < 1{
        return fmt.Errorf("ansi: the color should be between 1 and 255 you've passed %d, c)
	fmt.Printf("%s[38;5;%dm",ESC, c)
    return nil
}

