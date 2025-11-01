package main

import (
	"fmt"
	"time"

	"github.com/ams-tech-fin/tcs"
)

func main() {
	// Commodore 64 style terminal
	tcs.SetFullColors(tcs.BrightCyan, tcs.Blue)

	fmt.Println("\n    **** COMEDORME 64 GOLANG V1 ****")
	fmt.Println("\n 64K RAM SYSTEM  38911 BASIC BYTES FREE")
	fmt.Println("\n READY.")
	fmt.Println(" █")
	fmt.Println()

	time.Sleep(5 * time.Second)

	// Reset
	tcs.Reset()
	tcs.ClearScreen()
	fmt.Println("Demo finished!")
}
