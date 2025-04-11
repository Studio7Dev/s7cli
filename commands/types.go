package commands

import (
	"github.com/c-bata/go-prompt"
	. "mk3cli/s7cli/colors"
)

type Arg struct {
	Name     Name
	Datatype string
	Required bool
}

type Args []Arg

type Name struct {
	Full  string
	Short string
}

// Command struct

type Command struct {
	Name          string
	Description   string
	Args          Args
	SubCompletion []prompt.Suggest
	Exec          func(input []string, this Command) error
}

type CMDHandler struct {
	prompt      string
	buffer      []byte
	commands    []Command
	completions []prompt.Suggest
}

// Displays the valid usage of a command to the terminal
func (this Command) DisplayUsage() {
	usage := this.Name + " "

	for _, a := range this.Args {
		if a.Required {
			usage += "--" + a.Name.Format(false) + " [" + Cyan + a.Datatype + White + "] " + Reset
		} else {
			usage += Gray + "--" + a.Name.Format(true) + " [" + Cyan + a.Datatype + Gray + "] " + Reset
		}
	}
	println(usage)
}
