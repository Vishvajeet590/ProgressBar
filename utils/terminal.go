package utils

import (
	tsize "github.com/kopoli/go-terminal-size"
	"os"
	"os/exec"
	"runtime"
)

var Window tsize.Size

func GetTerminalSize() {
	w, err := tsize.GetSize()
	Window = w
	if err != nil {
		return
	}
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
