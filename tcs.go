//	║ ║ ║ ║ ║ ║ ║ ║ ║ ║ ║ ║
// ╔═══════════════════════╗
// ║       AMS-TCS         ║   Terminal Color Script
// ╚═══════════════════════╝   ©2025 - AMS TECNOLOGIA LTDA
//	║ ║ ║ ║ ║ ║ ║ ║ ║ ║ ║ ║

package tcs

import "fmt"

// Color represents available terminal colors
type Color int

const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	BrightBlack
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

// ANSI color codes
var ansiTextColors = map[Color]string{
	Black:         "\033[30m",
	Red:           "\033[31m",
	Green:         "\033[32m",
	Yellow:        "\033[33m",
	Blue:          "\033[34m",
	Magenta:       "\033[35m",
	Cyan:          "\033[36m",
	White:         "\033[37m",
	BrightBlack:   "\033[90m",
	BrightRed:     "\033[91m",
	BrightGreen:   "\033[92m",
	BrightYellow:  "\033[93m",
	BrightBlue:    "\033[94m",
	BrightMagenta: "\033[95m",
	BrightCyan:    "\033[96m",
	BrightWhite:   "\033[97m",
}

var ansiBackgroundColors = map[Color]string{
	Black:         "\033[40m",
	Red:           "\033[41m",
	Green:         "\033[42m",
	Yellow:        "\033[43m",
	Blue:          "\033[44m",
	Magenta:       "\033[45m",
	Cyan:          "\033[46m",
	White:         "\033[47m",
	BrightBlack:   "\033[100m",
	BrightRed:     "\033[101m",
	BrightGreen:   "\033[102m",
	BrightYellow:  "\033[103m",
	BrightBlue:    "\033[104m",
	BrightMagenta: "\033[105m",
	BrightCyan:    "\033[106m",
	BrightWhite:   "\033[107m",
}

const resetANSI = "\033[0m"

// SetTextColor sets the text color
func SetTextColor(color Color) {
	fmt.Print(ansiTextColors[color])
}

// SetBackgroundColor sets the background color
func SetBackgroundColor(color Color) {
	fmt.Print(ansiBackgroundColors[color])
}

// SetColors sets both text and background colors
func SetColors(textColor, bgColor Color) {
	fmt.Print(ansiBackgroundColors[bgColor] + ansiTextColors[textColor])
}

// ClearScreen clears the terminal screen
func ClearScreen() {
	fmt.Print("\033[2J\033[H")
}

// SetFullBackground fills the entire terminal with a background color
func SetFullBackground(bgColor Color) {
	fmt.Print(ansiBackgroundColors[bgColor])
	fmt.Print("\033[2J\033[H")
}

// SetFullColors fills the entire terminal with text and background colors
func SetFullColors(textColor, bgColor Color) {
	fmt.Print(ansiBackgroundColors[bgColor] + ansiTextColors[textColor])
	fmt.Print("\033[2J\033[H")
}

// Reset restores default terminal colors
func Reset() {
	fmt.Print(resetANSI)
}

// Printf prints formatted colored text
func Printf(textColor, bgColor Color, format string, args ...interface{}) {
	SetColors(textColor, bgColor)
	fmt.Printf(format, args...)
	Reset()
}

// Println prints a colored line
func Println(textColor, bgColor Color, text string) {
	SetColors(textColor, bgColor)
	fmt.Println(text)
	Reset()
}
