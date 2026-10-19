package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// websites list
	web := []string{"https://google.com", "https://youtube.com"}

	for _, url := range web {

		timerStart := time.Now()

		callback, err := http.Get(url)

		// error checker
		if err != nil {
			fmt.Printf("Site Down\n")
			continue
		}

		// close the call
		callback.Body.Close()

		timerEnd := time.Since(timerStart).Milliseconds()

		// Print out
		fmt.Println("Server -", url)
		fmt.Println("Status Code:", callback.StatusCode)
		fmt.Println("TIme taken:", timerEnd, "ms")
	}
}
