package main

import "fmt"

// START PAGE1 OMIT
type State int

const (
	Start State = iota
	Initializing
	Processing
	Finalizing
	Terminated
)

func (s State) GetState() State {
	return s
}

func (s *State) SetState(n State) {
	*s = n
}

// END PAGE1 OMIT

// START PAGE2 OMIT
func (s State) String() string {
	switch s {
	case Start:
		return "Start"
	case Initializing:
		return "Initializing"
	case Processing:
		return "Processing"
	case Finalizing:
		return "Finalizing"
	case Terminated:
		return "Terminated"
	default:
		return "Unknown"
	}
}

// END PAGE2 OMIT

// START PAGE3 OMIT
type Foo struct {
	State // 無名埋め込み // HL
	// state State とすると、Foo.stateとして明示的にアクセスする必要が
	// あるが、無名埋め込みならFoo.state.GetState()ではなく、自動的に 
	// Foo.GetState() が Stateへ委譲される。
}
type Bar struct{ State }
type Baz struct{ State }

func main() {
	f := Foo{}
	for i := 0; i < 5; i++ {
		fmt.Printf("Current state is: %s\n", f.GetState())
		f.SetState(State(int(f.GetState()) + 1))
	}
}

// END PAGE3 OMIT

// START PAGE4 OMIT
type Stater interface { // HL
	GetState() State  // HL
	SetState(s State) // HL
} // HL

// END PAGE4 OMIT

// START PAGE5 OMIT
func main2() {
	// Stater interfaceを満たすオブジェクト達なので、
	// Stater を保持する関数にまるっと入れることができる
	list := []Stater{Foo{}, Bar{}, Baz{}} // HL
	for _, s := range list {
		runStates(s)
	}
}

// Stater interfaceを満たしている構造体ならなんでもこの関数に突っ込める！
func runStates(s Stater) { // HL
	for i := 0; i < 5; i++ {
		fmt.Printf("Current state is: %s\n", s.GtState())
		s.SetState(State(int(s.GetState()) + 1))
	}
}

// END PAGE5 OMIT
