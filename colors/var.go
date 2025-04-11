package colors

type Color int

const (
	CBlack Color = iota
	CRed
	CGreen
	CYellow
	CBlue
	CMagenta
	CCyan
	CWhite
	None = 9
)

type Format int

const (
	// Format codes
	Default Format = iota
	Bold
	Dim
	Italic
	Underline
	Blinking
	Reverse
	Hidden
	Strikethrough
)

var (
	int2str = []string{
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}
)

// Builds an ANSI escape sequence
func Render(fg Color, bg Color, format ...Format) string {
	var formats string
	if len(format) > 0 {
		for _, f := range format {
			formats += ";" + int2str[f]
		}
	} else {
		formats = ""
	}
	return "\033[3" + int2str[fg] + ";4" + int2str[bg] + formats + "m"
}

// Creates a string using ANSI escape codes
func SRender(input string, fg Color, bg Color, format ...Format) string {
	var formats string
	if len(format) > 0 {
		for _, f := range format {
			formats += ";" + int2str[f]
		}
	} else {
		formats = ""
	}
	return "\033[3" + int2str[fg] + ";4" + int2str[bg] + formats + "m" + input + "\033[0m"
}

var (
	Reset   = Render(None, None, Default)
	Black   = Render(CBlack, None)
	Red     = Render(CRed, None)
	Green   = Render(CGreen, None)
	Yellow  = Render(CYellow, None)
	Blue    = Render(CBlue, None)
	Magenta = Render(CMagenta, None)
	Cyan    = Render(CCyan, None)
	White   = Render(CWhite, None)
	Gray    = Render(CWhite, None, Dim)
)
