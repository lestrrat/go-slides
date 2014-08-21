package main

import "fmt"

// START NOISE MAKER OMIT
type NoiseMaker interface {
	NoiseMake()
}

func main() {
	noiseMakers := []NoiseMaker{
		NewDog(),
		NewCat(),
		NewMouse(),
	}
	for _, v := range noiseMakers {
		v.NoiseMake()
	}
}
// END NOISE MAKER OMIT

type Dog struct{} // ちょっとよくないけど、とりあえず…

func (d Dog) NoiseMake() {
	fmt.Printf("%s goes '%s'\n", d.Type(), d.Sound())
}

func (d Dog) Type() string {
	return "Dog"
}

func (d Dog) Sound() string {
	return "Bow-wow"
}
