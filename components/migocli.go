// Package components is a set of structs used for creating
// interactive and command based CLIs.

// It covers basic components like select menus, checklists and questionaires
// and also you can create your own Interactive component by implementing the
// InteractivComponent interface.
package components

import "fmt"

// represents the cli application
type Cli struct {
	am         *AnsiManager
	components map[string]Component
}

// adds a component to the cli and if it needs
// the ansi manager assigns it to the component.
func (c *Cli) AddComponent(ca Component) {
	c.components[ca.GetName()] = ca
	if ca.NeedAnsi() {
		ca.SetAnsi(c.am)
	}
}

// calls the component with the name in cn
func (c *Cli) Start(cn string) {
	c.components[cn].Call()
}

// A command reads the flags and executes a function
// based on them, the function should be implemented
// in a way that uses the defined flags.
type Command struct {
	name     string
	flags    map[string]string
	callback func()
	next     *InteractiveComponent
}

// executes the callback function
func (c *Command) Call() {
	c.callback()
}

// ways to interact with a Command or Interactive component
type Component interface {
	GetName() string
	NeedAnsi() bool
	SetAnsi(a *AnsiManager) bool
	Call()
}

// ways to interact with an Interactive component
type InteractiveComponent interface {
	Component
	render() // renders the component in the terminal emulator
	interact()
}

// Menu to select one option
type SelectMenu struct {
	prompt     string
	indicator  string
	options    []string
	optionMap  map[int]string
	start_line int
	end_line   int
	next       *InteractiveComponent
	prev       *InteractiveComponent
	a          *AnsiManager
}

// Creates an interactive menu,
// p is the initial prompt
// i is the indicator if empty string it gives "❯" as default
// opts, are the selectable options of the menu
// next is a pointer to the next interactive component, it can be nil
// prev is a pointer to the previous interactive component
func NewSelectMenu(p, i string, opts []string, next, prev *InteractiveComponent) *SelectMenu {
	if len(i) == 0 {
		i = "❯"
	}
	i += " "
	return &SelectMenu{p, i, opts, make(map[int]string, len(opts)), 0, 0, next, prev, nil}
}

// Starts the Select menu
func (sm *SelectMenu) Call() {
	sm.render()
}

// Sets the AnsiManager for the Select menu
func (sm *SelectMenu) SetAnsi(a *AnsiManager) {
	sm.a = a
}

// renders the menu in the terminal emulator.
func (sm *SelectMenu) render() {
	sm.a.HideCursor()
	sm.a.WriteText(sm.prompt)

	locations := []int{}
	for _, op := range sm.options {
		sm.a.WriteText("  " + op)
		opLocation, _ := sm.a.GetCurrentLine()
		locations = append(locations, opLocation)
		sm.optionMap[opLocation] = op
	}
	sm.a.MoveCursor(locations[0])
	sm.a.WriteText(sm.indicator)
	// fmt.Printf("\033[38;5;46m%s\033[0m%s", indicator, m.optMap[locations[0]])
	sm.start_line = locations[0]
	sm.end_line = locations[len(locations)-1]
	sm.interact()

	sm.a.ShowCursor()
}

// sets an indicator in the selected line, and
// defines the way to move around the menu
func (sm *SelectMenu) interact() error {
	p := make([]byte, 1)
	c := sm.start_line
	// indicator := "❯ "
	for {
		sm.a.read(p)
		sm.a.RemoveText(c, 0, len(sm.indicator))
		switch string(p) {
		case "k":
			c--
			if c < sm.start_line {
				c = sm.end_line
			}
		case "j":
			c++
			if c > sm.end_line {
				c = sm.start_line
			}
		}
		sm.a.MoveCursor(c)
		sm.a.WriteText(sm.indicator)
		if string(p) == "q" {
			break
		}
	}
	option, ok := sm.optionMap[c]
	if !ok {
		return fmt.Errorf("Selected option not found")
	}
	fmt.Printf("Selected option: '%s'", option)
	return nil
}

type CheckList struct {
	prompt          string
	check_indicator string
	options         []string
	optionMap       map[int]string
	next            *InteractiveComponent
	prev            *InteractiveComponent
}

// shows the advancement on a set of tasks
type Loader struct {
	next *InteractiveComponent
	prev *InteractiveComponent
}

type Question struct {
	question string
	format   string
}

type Questionaire struct {
	prompt    string
	questions map[int]Question
	next      *InteractiveComponent
	prev      *InteractiveComponent
}
