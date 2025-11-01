# TCS - Terminal Color Script

Simple terminal color manipulation for Linux/Unix systems using ANSI escape codes.

[![Go Report Card](https://goreportcard.com/badge/github.com/ams-tech-fin/tcs)](https://goreportcard.com/report/github.com/ams-tech-fin/tcs)
[![GoDoc](https://godoc.org/github.com/ams-tech-fin/tcs?status.svg)](https://godoc.org/github.com/ams-tech-fin/tcs)
[![GitHub release](https://img.shields.io/github/release/SEUUSUARIO/termcolor.svg)](https://github.com/ams-tech-fin/tcs/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
## Features

- 🎨 16 colors (8 normal + 8 bright)
- 📝 Text color manipulation
- 🖼️ Background color manipulation
- 🖥️ Full terminal screen coloring
- 🧹 Clear screen
- 🔄 Reset to default colors
- 📦 Zero dependencies
- 🐧 Linux/Unix/Mac compatible

## Installation

```bash
go get github.com/ams-tech-fin/tcs
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/ams-tech-fin/tcs"
)

func main() {
    // Set text color
    tcs.SetTextColor(tcs.BrightRed)
    fmt.Println("Red text")
    
    // Set background color
    tcs.SetBackgroundColor(tcs.Blue)
    fmt.Println("Blue background")
    
    // Set both
    tcs.SetColors(tcs.BrightYellow, tcs.Blue)
    fmt.Println("Yellow on blue")
    
    // Reset colors
    tcs.Reset()
}
```

## Available Colors

```go
tcs.Black
tcs.Red
tcs.Green
tcs.Yellow
tcs.Blue
tcs.Magenta
tcs.Cyan
tcs.White
tcs.BrightBlack   // Gray
tcs.BrightRed
tcs.BrightGreen
tcs.BrightYellow
tcs.BrightBlue
tcs.BrightMagenta
tcs.BrightCyan
tcs.BrightWhite
```

## API Reference

### Basic Functions

#### SetTextColor(color Color)
Sets the text color.

```go
tcs.SetTextColor(tcs.BrightGreen)
fmt.Println("Green text")
```

#### SetBackgroundColor(color Color)
Sets the background color.

```go
tcs.SetBackgroundColor(tcs.Red)
fmt.Println("Red background")
```

#### SetColors(textColor, bgColor Color)
Sets both text and background colors.

```go
tcs.SetColors(tcs.White, tcs.Black)
fmt.Println("White on black")
```

### Full Screen Functions

#### SetFullBackground(bgColor Color)
Fills the entire terminal with a background color.

```go
tcs.SetFullBackground(tcs.Blue)
fmt.Println("Entire terminal is now blue")
```

#### SetFullColors(textColor, bgColor Color)
Fills the entire terminal with text and background colors.

```go
tcs.SetFullColors(tcs.BrightCyan, tcs.Blue)
fmt.Println("Commodore 64 style!")
```

### Utility Functions

#### ClearScreen()
Clears the terminal screen.

```go
tcs.ClearScreen()
```

#### Reset()
Resets colors to terminal defaults.

```go
tcs.Reset()
```

#### Printf(textColor, bgColor Color, format string, args ...interface{})
Prints formatted colored text and automatically resets colors.

```go
tcs.Printf(tcs.BrightGreen, tcs.Black, "Hello %s!\n", "World")
```

#### Println(textColor, bgColor Color, text string)
Prints a colored line and automatically resets colors.

```go
tcs.Println(tcs.BrightCyan, tcs.Black, "Colored line")
```

## Examples

### Example 1: Commodore 64 Screen

```go
package main

import (
    "fmt"
    "github.com/ams-tech-fin/tcs"
)

func main() {
    tcs.SetFullColors(tcs.BrightCyan, tcs.Blue)
	fmt.Println("\n    **** COMEDORME 64 GOLANG V1 ****")
	fmt.Println("\n 64K RAM SYSTEM  38911 BASIC BYTES FREE")
	fmt.Println("\n READY.")
	fmt.Println(" █")
	fmt.Println()
}
```

### Example 2: Retro Terminal

```go
package main

import (
    "fmt"
    "github.com/ams-tech-fin/tcs"
)

func main() {
    tcs.SetFullColors(tcs.BrightGreen, tcs.Black)
    fmt.Println("\n  RETRO TERMINAL v1.0")
    fmt.Println("  > System initialized")
    fmt.Println("  > Ready for commands")
}
```

### Example 3: Colored Logging

```go
package main

import (
    "fmt"
    "github.com/ams-tech-fin/tcs"
)

func LogInfo(msg string) {
    tcs.Printf(tcs.BrightCyan, tcs.Black, "[INFO] %s\n", msg)
}

func LogError(msg string) {
    tcs.Printf(tcs.BrightRed, tcs.Black, "[ERROR] %s\n", msg)
}

func LogSuccess(msg string) {
    tcs.Printf(tcs.BrightGreen, tcs.Black, "[OK] %s\n", msg)
}

func main() {
    LogInfo("Starting application...")
    LogSuccess("Connected to database")
    LogError("Failed to load config")
}
```

### Example 4: Color Demo

```go
package main

import (
    "fmt"
    "github.com/ams-tech-fin/tcs"
)

func main() {
    colors := []struct {
        color tcs.Color
        name  string
    }{
        {tcs.Black, "Black"},
        {tcs.Red, "Red"},
        {tcs.Green, "Green"},
        {tcs.Yellow, "Yellow"},
        {tcs.Blue, "Blue"},
        {tcs.Magenta, "Magenta"},
        {tcs.Cyan, "Cyan"},
        {tcs.White, "White"},
        {tcs.BrightRed, "BrightRed"},
        {tcs.BrightGreen, "BrightGreen"},
        {tcs.BrightYellow, "BrightYellow"},
        {tcs.BrightBlue, "BrightBlue"},
        {tcs.BrightMagenta, "BrightMagenta"},
        {tcs.BrightCyan, "BrightCyan"},
        {tcs.BrightWhite, "BrightWhite"},
    }

    fmt.Println("Available colors:")
    for _, c := range colors {
        tcs.SetTextColor(c.color)
        fmt.Printf("█ %-20s", c.name)
        tcs.SetBackgroundColor(c.color)
        fmt.Print("    ")
        tcs.Reset()
        fmt.Println()
    }
}
```

## Compatibility

- ✅ Linux
- ✅ macOS
- ✅ Unix-like systems

## License

MIT License - feel free to use in your projects!

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Author

Created with ❤️ for André "Dex" Mendonça e Silva