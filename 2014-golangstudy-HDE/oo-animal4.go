type Dog struct {
	Animal // Animalの要素、メソッドに関しては自動的に委譲する、という宣言
}

d := NewDog()

d.NoiseMake() // これと
d.Animal.NoiseMake() // これは全く同じ。Type()とSound()はAnimalの物が使われる
