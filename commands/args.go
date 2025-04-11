package commands

import (
	. "mk3cli/s7cli/colors"
)

func (this Name) Format(is_grayed bool) string {
	output := ""
	char_found := false

	format := []string{Reset + Gray, Render(CWhite, None, Underline)}
	if is_grayed {
		format = []string{Reset + Gray, Render(CWhite, None, Dim, Underline)}
	}

	for _, c := range this.Full {
		if string(c) == this.Short && !char_found {
			output += format[1] + string(c) + format[0]
			char_found = true
			continue
		}
		output += string(c)
	}

	return output
}

func DisplayArgs(args []Arg) string {
	output := ""

	for _, a := range args {
		if a.Required {
			output += Reset + "<" + White + a.Name.Format(false) + Reset + " (" + Cyan + a.Datatype + Reset + ")> "
		} else {
			output += Reset + "[" + Gray + a.Name.Format(true) + White + " (" + Cyan + a.Datatype + White + ")] "
		}
	}
	return output
}
