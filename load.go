package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Parse flags
	interview := flag.Bool("interview", false, "Boolean flag for running interview logic on target")
	target := flag.String("target", "", "Target URL for load testing")

	flag.Parse()

	// Ensure target URL is provided
	if *target == "" {
		fmt.Println("Error: you must specify a target URL using -target=URL")
		os.Exit(1)
	}

	fmt.Println("URL target is ", *target)

	// Notify user if interview logic is enabled
	if *interview {
		fmt.Println("You will apply interview logic to ", *target)
	}
}
