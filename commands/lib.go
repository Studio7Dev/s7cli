package commands

import (
	"github.com/c-bata/go-prompt"
	"strings"
)

func (this *CMDHandler) getCommandArgCompletions(c Command) []prompt.Suggest {
	output := []prompt.Suggest{}
	for _, arg := range c.Args {
		name := arg.Name
		if arg.Required == true {
			name += "*"
		}

		output = append(output, prompt.Suggest{
			Text:        name,
			Description: arg.Datatype,
		})
	}
	return output
}

func (this *CMDHandler) forceBestCompletion(completions []prompt.Suggest, buffer *prompt.Buffer) {
	// check however may characters have already been typed by the user
	preCompletedCharCount := len(buffer.Document().GetWordBeforeCursor())

	if len(completions) == 0 || preCompletedCharCount == len(completions[0].Text) {
		return
	}

	// modify the input buffer to include completed text
	buffer.InsertText(completions[0].Text[(preCompletedCharCount):], false, true)
}

func (this *CMDHandler) updateCompletions(d prompt.Document) []prompt.Suggest {
	words := strings.Fields(d.Text)
	suggestions := []prompt.Suggest{}

	if len(words) == 0 || (len(words) == 1 && (d.GetCharRelativeToCursor(0) != ' ')) {

		for _, c := range this.commands {
			suggestions = append(suggestions, prompt.Suggest{
				Text:        c.Name,
				Description: c.Description,
			})
		}
		goto ret
	} else if len(words) >= 1 && d.GetCharRelativeToCursor(0) == ' ' {

		for _, c := range this.commands {
			if strings.ToLower(words[0]) == strings.ToLower(c.Name) {
				if c.SubCompletion != nil {
					suggestions = append(suggestions, c.SubCompletion(strings.Join(words[1:], " "))...)
				} else {
					goto ret
				}
				break
			}
		}
	} else {
		return this.completions
	}

ret:
	this.completions = suggestions
	return this.completions
}
