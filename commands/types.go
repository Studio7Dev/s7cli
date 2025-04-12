package commands

import (
	"github.com/c-bata/go-prompt"
	. "mk3cli/s7cli/colors"
)

type Arg struct {
	Name     string
	Datatype string
	Required bool
}

type Args []Arg

// Command struct

type Command struct {
	Name          string
	Description   string
	Args          Args
	SubCompletion func(input string) []prompt.Suggest
	Exec          func(input []string, this *Command) error
}

type CMDHandler struct {
	prompt        string
	buffer        []byte
	commands      []Command
	completions   []prompt.Suggest
	prevWordIndex int
}

// Displays the valid usage of a command to the terminal
func (this Command) DisplayUsage() {
	usage := this.Name + " "

	for _, a := range this.Args {
		if a.Required {
			usage += "--" + a.Name + " [" + Cyan + a.Datatype + White + "] " + Reset
		} else {
			usage += Gray + "--" + a.Name + " [" + Cyan + a.Datatype + Gray + "] " + Reset
		}
	}
	println(usage)
}
