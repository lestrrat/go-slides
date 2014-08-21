package main

import (
	"fmt"
	"time"
)

// START CODE OMIT
func main() {
	defer fmt.Println("main exits")
	go func() {
		defer fmt.Println("Deferred function NOT fired!") // これは発動しません :( // HL

		fmt.Println("goroutine starts...")
		time.Sleep(5 * time.Second)
	}()

	// ↑ のgoroutineの終了を待たずにメイン関数が終了してしまうのでgoroutineも途中で停止する
	fmt.Println("Sleeping for 0.5 secs...")
	time.Sleep(500 * time.Millisecond)
}

// END CODE OMIT
