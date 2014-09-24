// MODIFIED FROM ORIGINAL peco CODE TO FIT PRESENTATION

package peco

import "github.com/nsf/termbox-go"

// START SCREEN OMIT
// Screen hides termbox from tne consuming code so that
// it can be swapped out for testing
type Screen interface {
	Clear(termbox.Attribute, termbox.Attribute) error
	Flush() error
	PollEvent() chan termbox.Event
	SetCell(int, int, rune, termbox.Attribute, termbox.Attribute)
	Size() (int, int)
}
// END SCREEN OMIT

// Termbox just hands out the processing to the termbox library
type Termbox struct{}

// START TERMBOX OMIT
func (t Termbox) Clear(fg, bg termbox.Attribute) error { // HL
	return termbox.Clear(fg, bg)
}

func (t Termbox) Flush() error { // HL
	return termbox.Flush()
}

func (t Termbox) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) { // HL
	termbox.SetCell(x, y, ch, fg, bg)
}

func (t Termbox) Size() (int, int) { // HL
	return termbox.Size()
}

func (t Termbox) PollEvent() chan termbox.Event { // HL
	evCh := make(chan termbox.Event)
	go func() {
		defer func() { recover() }()
		defer func() { close(evCh) }()
		for {
			evCh <- termbox.PollEvent()
		}
	}()
	return evCh
}
// END TERMBOX OMIT
