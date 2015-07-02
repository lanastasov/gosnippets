package main

import (
	"fmt"
	g "github.com/lanastasov/gosnippets/lib"
)

func main() {
	fmt.Println("----")
	fmt.Println(g.RepeatString("*", 8))
	fmt.Println(g.Pow(2, 12))
	// fmt.Println(g.Closure())
	fmt.Println(g.Add(1, 2, 3, 4))
}
