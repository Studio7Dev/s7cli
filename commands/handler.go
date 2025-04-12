package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/c-bata/go-prompt"
	. "mk3cli/s7cli/colors"
)

func NewHandler(prompt string, commands ...Command) CMDHandler {
	return CMDHandler{prompt: prompt, commands: commands}
}

func (this *CMDHandler) GetInput() string {
	return prompt.Input(
		// Prompt text
		this.prompt,

		// Completion engine
		func(d prompt.Document) []prompt.Suggest {
			if d.GetCharRelativeToCursor(0) == ' ' {
				this.updateCompletions(d)
				// TODO:
				// - cache the current word position to determine if we need to update the completion cache
			}

			return prompt.FilterHasPrefix(this.completions, d.GetWordBeforeCursor(), true)
		},

		// Completion styling options
		prompt.OptionSuggestionBGColor(prompt.DarkGray),
		prompt.OptionSuggestionTextColor(prompt.Black),
		prompt.OptionDescriptionBGColor(prompt.LightGray),
		prompt.OptionDescriptionTextColor(prompt.Black),

		prompt.OptionSelectedSuggestionBGColor(prompt.DarkGreen),
		prompt.OptionSelectedSuggestionTextColor(prompt.White),
		prompt.OptionSelectedDescriptionBGColor(prompt.Green),
		prompt.OptionSelectedDescriptionTextColor(prompt.White),

		// Add bind to ? key to show completions
		prompt.OptionAddASCIICodeBind(prompt.ASCIICodeBind{
			ASCIICode: []byte{0x3F},
			Fn: func(buffer *prompt.Buffer) {
				d := buffer.Document()
				compMatches := prompt.FilterHasPrefix(this.completions, d.GetWordBeforeCursor(), true)

				if len(compMatches) == 1 {
					this.forceBestCompletion(compMatches, buffer)
					return
				}

				fmt.Println("?\n\u001B[2K")
				for _, c := range compMatches {
					fmt.Println("\033[2K\r  " + c.Text + " " + SRender(c.Description, CWhite, None, Dim))
				}
				fmt.Println("\033[2K\r")
			},
		}),
		// Add bind to space key to auto complete the first suggestion
		prompt.OptionAddASCIICodeBind(prompt.ASCIICodeBind{
			ASCIICode: []byte{0x20},
			Fn: func(buffer *prompt.Buffer) {
				this.forceBestCompletion(prompt.FilterHasPrefix(this.completions, buffer.Document().GetWordBeforeCursor(), true), buffer)
				buffer.InsertText(" ", false, true) // maintain normal space key behavior
				this.updateCompletions(*buffer.Document())
			},
		}),
		// Add bind to Ctrl+C to exit the program normally
		prompt.OptionAddKeyBind(prompt.KeyBind{
			Key: prompt.ControlC,
			Fn: func(buffer *prompt.Buffer) {
				fmt.Println(SRender("Goodbye 👋", CGreen, None, Bold))
				os.Exit(0)
			},
		}))
}

func (this CMDHandler) Handle(input string) {
	args := strings.Split(input, " ")
	notfound := true
	matches := []Command{}

	for _, c := range this.commands {
		if args[0] == c.Name {
			matches = append(matches, c)
			notfound = false
		}
	}

	if len(matches) == 1 {
		matches[0].Exec(args, &matches[0])
	}

	if notfound {
		fmt.Println(SRender("Command not found: '"+args[0]+"'", CRed, None))
	}
}

func (this *CMDHandler) SetPrompt(prompt string) {
	this.prompt = prompt
}

func (this *CMDHandler) AddCommand(command Command) {
	// check to verify required args come before any non-required
	in_required := true
	for _, arg := range command.Args {
		if !arg.Required && in_required {
			in_required = false
		} else if arg.Required && !in_required {
			panic("Command \"" + command.Name + "\" has required argument after non-required arguments!\n\tArgument: " + arg.Name)
			os.Exit(1)
		}
	}
	this.commands = append(this.commands, command)
	this.completions = append(this.completions, prompt.Suggest{
		Text:        command.Name,
		Description: command.Description,
	})
}

// initalizes the CLI with 3 default commands (help, clear, exit)
func (this *CMDHandler) Init() CMDHandler {
	// Add the default commands

	// Exit command
	this.AddCommand(Command{
		Name:          "exit",
		Description:   "Exits this application.",
		Args:          []Arg{},
		SubCompletion: nil,
		Exec: func(args []string, command *Command) error {
			fmt.Println(SRender("Goodbye 👋", CGreen, None, Bold))
			os.Exit(0)
			return nil
		},
	})

	// Help command
	this.AddCommand(Command{
		Name:          "help",
		Description:   "Displays the current list of commands",
		Args:          []Arg{},
		SubCompletion: nil,
		Exec: func(args []string, command *Command) error {
			if len(args) > 1 {
				for _, c := range this.commands {
					if args[0] == c.Name {
						c.DisplayUsage()
						println()
					}
				}
			}
			fmt.Println("List of all currently supported commands:\n")
			for _, c := range this.commands {
				print("  ")
				c.DisplayUsage()
			}
			return nil
		},
	})

	// Clear command
	this.AddCommand(Command{
		Name:          "clear",
		Description:   "Clears the console.",
		Args:          []Arg{},
		SubCompletion: nil,
		Exec: func(input []string, this *Command) error {
			os_switch := make(map[string]func()) //Initialize it
			os_switch["linux"] = func() {
				cmd := exec.Command("clear") //Linux example, its tested
				cmd.Stdout = os.Stdout
				cmd.Run()
			}
			os_switch["windows"] = func() {
				cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
				cmd.Stdout = os.Stdout
				cmd.Run()
			}

			value, ok := os_switch[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
			if ok {                              //if we defined a clear func for that platform:
				value() //we execute it
			} else { //unsupported platform
				fmt.Println("Failed; Your terminal isn't ANSI! :(")
			}
			return nil
		},
	})
	return *this
}
