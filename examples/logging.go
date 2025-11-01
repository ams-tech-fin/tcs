package main

import (
	"fmt"
	"github.com/ams-tech-fin/tcs"
	"time"
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

func LogWarning(msg string) {
	tcs.Printf(tcs.BrightYellow, tcs.Black, "[WARN] %s\n", msg)
}

func main() {
	fmt.Println("=== Colored Logging Example ===\n")

	LogInfo("Starting application...")
	time.Sleep(500 * time.Millisecond)

	LogInfo("Connecting to database...")
	time.Sleep(500 * time.Millisecond)

	LogSuccess("Connected to database")
	time.Sleep(500 * time.Millisecond)

	LogInfo("Loading configuration...")
	time.Sleep(500 * time.Millisecond)

	LogWarning("Configuration file not found, using defaults")
	time.Sleep(500 * time.Millisecond)

	LogInfo("Starting web server...")
	time.Sleep(500 * time.Millisecond)

	LogSuccess("Server started on port 8080")
	time.Sleep(500 * time.Millisecond)

	LogError("Failed to load plugin: plugin_xyz.so")
	time.Sleep(500 * time.Millisecond)

	LogInfo("Application ready!")
}
