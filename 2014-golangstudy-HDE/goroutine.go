package main

import (
	"fmt"
	"sync"
	"time"
)

func doLongRunningProcess(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// ... do something ...
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go doLongRunningProcess(wg)

	wg.Wait()
}

func bailoutFromGoroutine() {
	go func() {
		return // HL
	}()
}

func noDeferGoroutine() {
	go func() {
		defer func() {
			fmt.Println("Defered function") // 永遠に呼ばれない… // HL
		}()

		for { // 無限ループ // HL
			// 色々やる
		}
	}()
}

func quitChanGoroutine(wg *sync.WaitGroup, quit chan struct{}) {
	go func() {
		defer wg.Done() // quitにお知らせがあれば、ちゃんと呼ばれる // HL

		for { // 無限ループ // HL
			select {
			case <-quit: // 終了のお知らせ！ // HL
				return // BAIL OUT // HL
			default:
				// 色々やる
			}
		}
	}()
}

func startStopGoroutines() {
	wg := &sync.WaitGroup{}
	quit := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		quitChanGoroutine(wg, quit)
	}

	// 10秒後に終了！
	time.AfterFunc(10*time.Second, func() {
		close(quit)
	})

	wg.Wait()
}
