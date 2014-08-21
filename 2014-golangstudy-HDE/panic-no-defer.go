package main

import "fmt"

func main() {
	defer func() { // WILL NOT
		fmt.Println("Defered function called!")
	}()

	panic("Panic!")
}
