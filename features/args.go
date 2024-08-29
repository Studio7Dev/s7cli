package features

import (
	"fmt"
	. "mk3cli/s7cli/commands"
)

func PrintEnabledArgs(enabled FeatureSet) {
	output := "\n   ┌ " + Green + enabled.Feature.Name + White + " "
	args := enabled.Feature.Args

	x := 0
	for _, a := range args {
		enabledVal := fmt.Sprintf("%v", enabled.Args[x].Value)

		if len(enabledVal) > 10 {
			enabledVal = enabledVal[0:10] + Gray + "..."
		}

		if a.Required {
			output += "<" + a.Name.Full + "=" + fmt.Sprintf("\""+Green+"%v"+White+"\"", enabledVal) + " (" + Gray + a.Datatype + White + ")> "
		} else if !a.Required && x < len(enabled.Args) {
			output += "[" + a.Name.Full + "=" + fmt.Sprintf("\""+Green+"%v"+White+"\"", enabledVal) + " (" + Gray + a.Datatype + White + ")] "
		} else {
			output += "[" + a.Name.Full + " (" + Gray + a.Datatype + White + ")] "
		}
		x++
	}

	output += "\n   └───> " + Gray + enabled.Feature.Description + Reset

	fmt.Println(output)
}
