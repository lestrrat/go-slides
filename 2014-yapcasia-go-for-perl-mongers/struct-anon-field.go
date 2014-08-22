package main

type Bar struct{}

func (b Bar) BarMethod() {}

// START CODE OMIT
type Foo struct {
	Bar // BarをFooに無名埋め込み // HL
}

func main() {
	f := Foo{}
	f.BarMethod() // 自動的に移譲！ // HL
}

// END CODE OMIT
