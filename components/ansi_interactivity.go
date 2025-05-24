package components

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

const (
	ESC           = "\033"
	HIDE_CURSOR   = "[?25l"
	SHOW_CURSOR   = "[?25h"
	LOCATE_CURSOR = "[6n"
	NEXT_LINE     = "[1E"
	RESET         = "[0m"
)

// Manages the terminal emulator through ansi escape codes.
type AnsiManager struct {
	c *bufio.ReadWriter
}

func (a *AnsiManager) write(code string) {
	a.c.Write([]byte(code))
	a.c.Flush()
}

// moves cursor to the line in position l
func (a *AnsiManager) MoveCursor(l int) {
	code := fmt.Sprintf("%s[%d;0H", ESC, l)
	a.write(code)
}

// sets the color of the character's from that point on
func (a *AnsiManager) SetForegroundColor(color int) error {
	if color > 255 || color < 1 {
		return fmt.Errorf("ansi: the color should be between 1 and 255, you've passed %d", color)
	}

	code := fmt.Sprintf("%s[38;5;%dm", ESC, color)
	a.write(code)
	return nil
}

// returns the current cursor location in terms of
// a line and a column
func (a *AnsiManager) GetCurrentLine() (int, int) {
	code := fmt.Sprintf("%s%s", ESC, LOCATE_CURSOR)
	a.write(code)
	v := ""
	s := ""
	p := make([]byte, 1)
	for s != "R" {
		a.c.Read(p)
		s = string(p)
		v += s
	}
	// regex for capturing the current line in the terminal
	rp := regexp.MustCompile(`(\d{1,2})`)
	mp := rp.FindStringSubmatch(v)
	l, _ := strconv.Atoi(mp[0])
	c, _ := strconv.Atoi(mp[1])

	return l, c
}

// restores the formatting when its called
func (a *AnsiManager) ClearFormatting() {
	a.write(fmt.Sprintf("%s%s", ESC, RESET))
}

// writes text on a new line
func (a *AnsiManager) WriteText(t string) {
	a.write(fmt.Sprintf("%s%s%s", ESC, NEXT_LINE, t))
}

// writes text at the beginning of a given line
func (a *AnsiManager) WriteTextOn(t string, line int) {
	a.MoveCursor(line)
	a.write(t)
}

// writes text with the given color and resets
// the formating back
func (a *AnsiManager) WriteTextColored(t string, c int) {
	code := fmt.Sprintf("%s[38;5;%dm%s", ESC, c, t)
	a.write(code)
}

// writes text with a given 256 color in a given line
func (a *AnsiManager) WriteTextColoredOn(t string, c, l int) {
	a.MoveCursor(l)
	code := fmt.Sprintf("%s[38;5;%dm%s", ESC, c, t)
	a.write(code)
	a.ClearFormatting()
}

// removes i characters at line l, from column c
// the cursor remains in its current position
func (a *AnsiManager) RemoveText(l, c, i int) {
	code := fmt.Sprintf("ESC[%d;%d;%dH", l, c, i)
	a.write(code)
}
