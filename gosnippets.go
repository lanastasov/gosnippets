package gosnippets
import (
		"os"
		"os/exec"
)

func ScreenCls() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
}
