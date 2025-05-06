package components

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type Menu struct {
	console *bufio.ReadWriter
	prompt  string
	options []string
	optMap  map[int]string
}

func NewMenu(console *bufio.ReadWriter, prompt string, options []string) *Menu {
	return &Menu{
		console, prompt, options,
		make(map[int]string, len(options)),
	}
}

func moveCursorToLine(line int) {
	fmt.Printf("\033[%d;0H", line)
}

// Renders the menu in the command line
func (m *Menu) Render() {
	fmt.Print("\033[?25l")
	m.console.Write([]byte("\033[1E" + m.prompt))
    m.console.Flush()
	locations := []int{}
	for _, op := range m.options {
		fmt.Printf("\033[1E  %s", op)
		opLocation := m.getCurrentLine()
		locations = append(locations, opLocation)
		m.optMap[opLocation] = op
	}
	moveCursorToLine(locations[0])
	indicator := "❯ "
	fmt.Printf("\033[38;5;46m%s\033[0m%s", indicator, m.optMap[locations[0]])

	m.startOptionSelect(locations[0], locations[len(locations)-1])

	fmt.Print("\033[?25h")
}

func (m *Menu) startOptionSelect(s, e int) error {
	p := make([]byte, 1)
	c := s
	indicator := "❯ "
	for {
		m.console.Read(p)
		switch string(p) {
		case "k":
			moveCursorToLine(c)
			fmt.Printf("  %s  ", m.optMap[c])
			c--
			if c < s {
				c = e
			}
		case "j":
			moveCursorToLine(c)
			fmt.Printf("  %s  ", m.optMap[c])
			c++
			if c > e {
				c = s
			}
		}
		moveCursorToLine(c)
		fmt.Printf("\033[38;5;46m%s\033[0m%s", indicator, m.optMap[c])
		if string(p) == "q" {
			break
		}
	}
	option, ok := m.optMap[c]
	if !ok {
		return fmt.Errorf("Selected option not found")
	}
	fmt.Printf("Selected option: '%s'", option)
	return nil
}

func (m *Menu) getCurrentLine() int {
	fmt.Print("\033[6n")
	v := ""
	s := ""
	p := make([]byte, 1)
	for s != "R" {
		m.console.Read(p)
		s = string(p)
		v += s
	}
	// regex for capturing the current line in the console
	rp := regexp.MustCompile(`(\d{1,2})`)
	mp := rp.FindStringSubmatch(v)
	l, _ := strconv.Atoi(mp[0])

	return l
}
