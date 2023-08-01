package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
)

// url
var (
	//go:embed launch_url
	url string
)

func main() {

	// browser
	chrome := []string{
		"C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe",
		"C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe",
		"C:\\Program Files\\Microsoft\\Edge\\Application\\msedge.exe",
		"C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe",
		"C:\\Program Files\\Mozilla Firefox\\firefox.exe",
		"C:\\Program Files (x86)\\Mozilla Firefox\\firefox.exe",
	}

	started := false

	fmt.Println("Launching application...")

	for _, app := range chrome {
		_, err := os.Open(app)
		if err == nil {
			err := exec.Command(app, url).Start()

			if err == nil {
				started = true
				break
			}
		}
	}

	if !started {
		fmt.Println("Please install a modern browser (Chrome, Edge, or Firefox) and try again...")
		fmt.Println("")
		fmt.Println("Press Enter to exit")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		fmt.Println()
	}
}
