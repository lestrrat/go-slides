package main

import "os/exec"

func main() {
	cmd := exec.Command("時間のかかるコマンド")
	go cmd.Run()

	// 本当はここでgoroutineを待たないといけないけど、
	// ここでは割愛
}
