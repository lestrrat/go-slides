package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		x := 0
		ticker := time.NewTicker(time.Second)
		for _ = range ticker.C {
			x++
			fmt.Printf("%d second elapsed...\n", x)
		}
	}()

// START CODE OMIT
	done := make(chan struct{})
	go func() {
		defer func() {
			done <- struct{}{} // このgoroutineが終了したらお知らせを送る // HL
		}()
		time.Sleep(5 * time.Second)
		fmt.Println("Goroutine fired!")
	}()

	<-done // お知らせが来るまで待つ // HL
// END CODE OMIT
}
