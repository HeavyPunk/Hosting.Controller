package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("java", "-jar", "/home/blackpoint/Games/minecraft/TLauncher-2.876.jar")
	cmd.Start()

	fmt.Print(cmd.Process.Pid)
}
