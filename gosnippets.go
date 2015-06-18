package gosnippets

import (
	"os"
	"os/exec"
	// "path"
	"fmt"
)

func ScreenCls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Watcher() {
	// _, filename, _, _ := runtime.Caller(1)
	// // path.Base(filename)
	// var watch = "go run" + path.Base(filename)
	// fmt.Print("->%s",watch)
	// // cmd := exec.Command("filewatcher", ".", "go run hello.go")
	// // cmd.Run()
	// // return filename
}

// return 0, and not 100 ,
// as 0 is the default value for an integer
func Defer() {

	aValue := new(int)

	defer fmt.Println(*aValue)

	for i := 0; i < 100; i++ {
		*aValue++
	}
}
