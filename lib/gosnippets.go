package gosnippets

import (
	"os"
	"os/exec"
	// "path"
	"bytes"
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

func RepeatString(st string, n int) string {
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		buffer.WriteString(st)
	}
	return buffer.String()
}

func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

func Add(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func Closure() int {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() { fmt.Println("defer_closure i = ", i) }()
		fs[i] = func() { fmt.Println("closure i = ", i) }
	}

	for _, f := range fs {
		a := 0
		if a > 0 {
			f()
		}
	}

	return 0
}
