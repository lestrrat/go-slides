type Dog struct {
	Animal
}

func NewAnimal() Dog {
	return Animal{}
}

func (d Dog) Type() string {
	return "Dog"
}

func (d Dog) Sound() string {
	return "Bow-wow"
}

d := NewDog()
d.NoiseMake()
