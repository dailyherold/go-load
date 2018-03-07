package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	// Parse flags
	interview := flag.Bool("interview", false, "Boolean flag for running interview logic on target URL")
	target := flag.String("target", "", "Target URL for load testing")
	loops := flag.Int("loops", 5, "Number of times you want to hit target URL with a GET request")

	flag.Parse()

	// Ensure target URL is provided
	if *target == "" {
		fmt.Println("Error: you must specify a target URL using -target=URL")
		os.Exit(1)
	}

	load(*target, *interview, *loops)
}

func load(url string, interview bool, loops int) {
	var output string
	delay := 1

	// limit loops to 3 for interview logic to satisfy 1s, 2s, 4s GET request timeout requirement
	if interview {
		loops = 3
	}

	for i := 1; i <= loops; i++ {
		// Set http client timeout
		client := http.Client{
			Timeout: timeout(delay),
		}

		// Log details about request
		output = fmt.Sprintf("target=%s, timeout=%ds, loop=%d, interview=%t", url, delay, i, interview)
		fmt.Println(output)

		// GET request
		resp, err := client.Get(url)

		// Log response
		if err != nil {
			fmt.Println(err)
			// If error, increase timeout by factor of two
			delay = delay * 2
		} else {
			fmt.Println(resp)
		}
	}
}

func timeout(seconds int) time.Duration {
	return time.Duration(seconds) * time.Second
}
