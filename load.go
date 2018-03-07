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

	load(*target, *interview)
}

func load(url string, interview bool) {
	output := fmt.Sprintf("Load func, url=%s, interview=%t", url, interview)
	fmt.Println(output)

	delay := 1

	for i := 1; i <= 3; i++ {
		client := http.Client{
			Timeout: timeout(delay),
		}

		fmt.Println("Timeout of ", delay)
		resp, err := client.Get(url)

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
