package features

import (
	"fmt"
	. "mk3cli/s7cli/colors"
	. "mk3cli/s7cli/commands"
)

func DisplayEnabledArgs(args []Arg, enabled []FeatureSetArg) string {
	output := " "

	x := 0
	for _, a := range args {
		enabledVal := fmt.Sprintf("%v", enabled[x].Value)

		if len(enabledVal) > 10 {
			enabledVal = enabledVal[0:10] + SRender("...", CWhite, None, Dim)
		}

		if a.Required {
			output += "<" + a.Name + "=" + fmt.Sprintf("\""+SRender("%v", CGreen, None)+"\"", enabledVal) + " (" + SRender(a.Datatype, CWhite, None, Dim) + ")> "
		} else if !a.Required && x < len(enabled) {
			output += "[" + a.Name + "=" + fmt.Sprintf("\""+SRender("%v", CGreen, None)+"\"", enabledVal) + " (" + SRender(a.Datatype, CWhite, None, Dim) + ")] "
		} else {
			output += "[" + a.Name + " (" + SRender(a.Datatype, CWhite, None, Dim) + ")] "
		}
		x++
	}
	return output
}
