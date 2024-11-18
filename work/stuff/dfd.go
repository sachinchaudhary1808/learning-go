package main

import (
	"fmt"
	"io/ioutil" // staticcheck: deprecated package
	"sync"
)

func main() {
	// Suspicious mutex usage
	var mu sync.Mutex
	mu.Lock()
	// Missing Unlock() - potential deadlock

	// Unnecessary conversion
	var x int = 42
	y := float64(int(x))

	// Unreachable code after return
	if x > 0 {
		return
	}
	fmt.Println("This will never be reached")

	// Using deprecated function
	data, _ := ioutil.ReadFile("test.txt") // error ignored and deprecated function
	fmt.Println(string(data))

	// Inefficient string comparison
	s := "hello"
	if len(s) > 0 && s[0] == 'h' && s[1] == 'e' && s[2] == 'l' {
		// Should use strings.HasPrefix instead
	}

	// Channel misuse
	ch := make(chan int)
	ch <- 1 // deadlock: sending to unbuffered channel with no receiver

	// Unused variable with potential side effect
	_ = fmt.Sprintf("unused %d", x)
}

// Function that could be simplified
func unnecessaryIf(x int) bool {
	if x > 0 {
		return true
	} else {
		return false
	}
}

// Context missing in HTTP client
func makeRequest() {
	// Missing context in http.Client
	http.Get("https://example.com")
}
