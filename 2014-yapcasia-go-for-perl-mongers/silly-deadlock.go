package main

func NewEvent() int {
	return 1
}

// START DEADLOCK LOOP OMIT
func deadlockLoop() {
	eventCh := make(chan int)
	for {
		event := <-eventCh
		switch event {
		case 1:
			// ...
		case 2:
			//...
		case 3:
			eventCh <- NewEvent() // oh shit // HL
		}
	}
}

// END DEADLOCK LOOP OMIT

// START DEADLOCK WORKAROUND OMIT
func deadlockLoop() {
	eventCh := make(chan int)
	for {
		event := <-eventCh
		switch event {
		case 1:
			// ...
		case 2:
			//...
		case 3:
			go func() { // HL
				eventCh <- NewEvent() // あまりよくないが、デッドロックは回避 // HL
			}() // HL
		}
	}
}

// END DEADLOCK WORKAROUND OMIT
