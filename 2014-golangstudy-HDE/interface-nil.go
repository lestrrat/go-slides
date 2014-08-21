package main

import "fmt"

// START SETUP OMIT
type MyInterface interface {
	Foo()
}

type MyConcreteType struct{}

func (m MyConcreteType) Foo() {}

func GetInterface() MyInterface {
	var m MyConcreteType
	return m
}
// END SETUP OMIT

// START MAIN OMIT
func main() {
	m := GetInterface()
	if m == nil {
		fmt.Println("m is nil")
	} else {
		fmt.Println("m is NOT nil")
	}
}
// END MAIN OMIT
