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
	loops := flag.Int("loops", 5, "Number of times you want to hit target URL with a GET request")

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

	load(*target, *interview, *loops)
}

func load(url string, interview bool, loops int) {
	output := fmt.Sprintf("Load func, url=%s, interview=%t", url, interview)
	fmt.Println(output)

	delay := 1

	if interview {
		loops = 3
	}

	for i := 1; i <= loops; i++ {
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
