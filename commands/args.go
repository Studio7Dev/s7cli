package commands

import (
	. "mk3cli/s7cli/colors"
)

func DisplayArgs(args []Arg) string {
	output := ""

	for _, a := range args {
		if a.Required {
			output += Reset + SRender(a.Name, CWhite, None, Bold) + ":" + SRender(a.Datatype, CCyan, None) + " "
		} else {
			output += Reset + SRender("[ ", None, None, Dim) + SRender(a.Name, CWhite, None) + ":" + SRender(a.Datatype, CCyan, None) + SRender("] ", None, None, Dim)
		}
	}
	return output
}
