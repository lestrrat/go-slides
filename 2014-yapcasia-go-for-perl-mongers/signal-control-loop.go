package main

import (
	"os"
	"os/signal"
	"syscall"
)

// START MAIN OMIT
func main() {
	done := make(chan struct{})
	quit := make(chan struct{})
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM) // シグナルを待つ // HL

	go doLoop(quit, done)
	go func() {
		<-sig              // シグナルを受け取ったので… // HL
		quit <- struct{}{} // チャンネルにお知らせ // HL
	}()

	<-done
}

// END MAIN OMIT
// START CLOSE SAMPLE OMIT
func main() {
	done := make(chan struct{})
	quit := make(chan struct{})
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM) // シグナルを待つ

	go doLoop(quit, done)
	go func() {
		<-sig       // シグナルを受け取ったので… 
		close(quit) // チャンネルにお知らせ // HL
	}()

	<-done
}

// END CLOSE SAMPLE OMIT

// START LOOPFUNC OMIT
func doLoop(quit, done chan struct{}) {
	defer func() { done <- struct{}{} }()

	loop := true
	go func() { // あくまで一例。selectを使う方法ももちろんある
		<-quit       // quitが読み込み可能になったら… // HL
		loop = false // ループ終了 // HL
	}()
	for loop {
		// 本来必要な処理
	}
}

// END LOOPFUNC OMIT
