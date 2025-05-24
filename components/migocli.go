// Package components is a set of structs used for creating
// interactive and command based CLIs.

// It covers basic components like select menus, checklists and questionaires
// and also you can create your own Interactive component by implementing the
// InteractivComponent interface.
package components

// represents the cli application
type Cli struct {
	am       AnsiManager
	commands map[string]Command 
	menus    map[string]Menu
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

// An interactive component is used by the CLI
type InteractiveComponent interface {
	render(a AnsiManager) // renders the component in the terminal emulator
	interact(a AnsiManager)
}

type SelectMenu struct {
	options   []string
	optionMap map[int]string
	next      *InteractiveComponent
	prev      *InteractiveComponent
}

type CheckList struct {
	options   []string
	optionMap map[int]string
	next      *InteractiveComponent
	prev      *InteractiveComponent
}

type Loader struct {
	next *InteractiveComponent
	prev *InteractiveComponent
}

type Question struct {
	question string
	format   string
}

type Questionaire struct {
	questions map[int]Question
	next      *InteractiveComponent
	prev      *InteractiveComponent
}
