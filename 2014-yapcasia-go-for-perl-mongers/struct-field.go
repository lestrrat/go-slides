package main

// START CODE OMIT
type Bar struct{}

func (b Bar) BarMethod() {}

type Foo struct {
	barField Bar // BarをFooに埋め込み // HL
}

func main() {
	f := Foo{}
	f.barField.BarMethod() // barFieldを明示的に呼び出す // HL
}

// END CODE OMIT
