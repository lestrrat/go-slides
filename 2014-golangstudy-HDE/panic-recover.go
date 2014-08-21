package main

import (
	"fmt"
	"os"
)

func init() {
	if _, err := os.Stat("/really/important/file"); err != nil {
		panic("/really/important/file does not exist!") // HL
	}
}

func recoverGoroutine() {
	go func() {
		defer func() { recover() }() // HL

		// ... code that might panic ...
	}()
}

