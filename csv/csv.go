package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// http://stackoverflow.com/questions/24555004/how-to-map-text-file-content-in-array-of-structure-in-go
func main() {

	pth := "content.txt"

	// What is this code doing?
	//↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
	type ChangeStatus struct {
		Nr1      string
		Nr2      string
		Category string
	}

	status := []ChangeStatus{}
	f, err := os.Open(pth)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer f.Close()

	r := csv.NewReader(f)
	for {
		parts, err := r.Read()
		if err == nil {
			cs := ChangeStatus{parts[0], parts[1], parts[2]}
			status = append(status, cs)
		} else {
			break
		}
	}
	fmt.Printf("%+v\n", status)

	//↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑

}

//content.txt
//a ,file!
//a ,line.

// OUTPUT:
/*
panic: runtime error: index out of range

goroutine 1 [running]:
main.main()
        C:/dev/go-lang/src/github.com/lanastasov/gosnippets/csv/csv.go:35 +0x3d1

goroutine 4 [runnable]:
runtime.runfinq()
        c:/go/src/runtime/malloc.go:712
runtime.goexit()
        c:/go/src/runtime/asm_amd64.s:2232 +0x1
exit status 2
*/
