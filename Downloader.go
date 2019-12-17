package main

import (
	"fmt"
	"os"
)

const argc int = 1

func main() {
	fmt.Println("Http Downloader Application")

	argv := os.Args[1:] //Arguments except the executable name, i.e excluding the 0th argument

	if len(argv) < argc {
		fmt.Printf("Total arguments shouold not be less than %d\n", argc)
		os.Exit(2)
	}
	url := argv[0]

	rc := Download(&url)
	if !rc {
		fmt.Println("Download failed")
		os.Exit(2)
	}
}
