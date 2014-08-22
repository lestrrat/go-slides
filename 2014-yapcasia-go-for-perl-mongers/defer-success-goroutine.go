package main

import (
	"fmt"
	"time"
)

// START CODE OMIT
func main() {
	defer fmt.Println("main exits")

	done := make(chan struct{})
	go func() {
		defer fmt.Println("Deferred function fired!") // これは発動します！ // HL
		defer func() { done <- struct{}{} }() // これを送るまでメイン関数を終了させない // HL

		fmt.Println("goroutine starts...")
		time.Sleep(time.Second)
	}()

	<-done // ↑のgoroutineをしっかり待つ！ // HL
}

// END CODE OMIT
