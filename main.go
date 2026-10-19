package test

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// time starts
	timerStart := time.Now()

	// call the website
	fmt.Println("Server - goole-test")
	callback, err := http.Get("https://google.com")

	// error checker
	if err != nil {
		fmt.Printf("Site Down\n")
		return
	}

	// close the call
	callback.Body.Close()

	// timer ends
	timerEnd := time.Since(timerStart).Milliseconds()

	fmt.Println("Status Code:", callback.StatusCode)
	fmt.Println("TIme taken:", timerEnd, "ms")
}
