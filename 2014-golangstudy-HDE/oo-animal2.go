import "fmt"
// START OO ANIMAL 2 OMIT
type Animal struct{}

func (a Animal) NoiseMake() {
	fmt.Printf("%s goes '%s'\n", a.Type(), a.Sound())
}

func (a Animal) Type() string {
	// Abstract
	return ""
}

func (a Animal) Sound() string {
	// Abstract
	return ""
}
// END OO ANIMAL 2 OMIT

//START OO ANIMAL 3 OMIT
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

func RunDog() {
	d := NewDog()
	d.NoiseMake()
}
//END OO ANIMAL 3 OMIT

