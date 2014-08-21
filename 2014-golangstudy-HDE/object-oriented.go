package main

import (
	"log"
	"reflect"
)

// START NOINHERITANCE OMIT
type Base struct {
}
type Child struct {
	Base
}

func (b Base) printReceiver() {
	log.Printf("I am: %s", reflect.TypeOf(b).Name())
}

func main() {
	a := Base{}
	b := Child{Base{}}

	a.printReceiver() // I am: Base // HL
	b.printReceiver() // I am: Base <--- あれ？ // HL
}
// END NOINHERITANCE OMIT

// START DELEGATION OMIT
func deleation() {
	child := Child{Base{}}
	child.printReceiver() // HL
	child.Base.printReceiver() // HL
}
// END DELEGATION OMIT
