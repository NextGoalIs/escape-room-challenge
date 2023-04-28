package utils

import (
	"os"
	"os/exec"
)

func ClearConsoleWindows() {
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}