package main

import (
	"fmt"
	"github.com/ams-tech-fin/tcs"
)

func main() {
	fmt.Println("=== Basic TCS Example ===\n")

	// Text colors
	tcs.SetTextColor(tcs.BrightRed)
	fmt.Println("Red text")
	tcs.SetTextColor(tcs.BrightGreen)
	fmt.Println("Green text")
	tcs.SetTextColor(tcs.BrightBlue)
	fmt.Println("Blue text")
	tcs.Reset()

	fmt.Println()

	// Background colors
	tcs.SetBackgroundColor(tcs.Red)
	fmt.Println("Red background")
	tcs.SetBackgroundColor(tcs.Green)
	fmt.Println("Green background")
	tcs.SetBackgroundColor(tcs.Blue)
	fmt.Println("Blue background")
	tcs.Reset()

	fmt.Println()

	// Combined colors
	tcs.SetColors(tcs.BrightYellow, tcs.Blue)
	fmt.Println("Yellow on blue")
	tcs.SetColors(tcs.White, tcs.Red)
	fmt.Println("White on red")
	tcs.Reset()

	fmt.Println()

	// Using helper functions
	tcs.Printf(tcs.BrightGreen, tcs.Black, "Hello %s!\n", "World")
	tcs.Println(tcs.BrightCyan, tcs.Black, "Colored line")
}
