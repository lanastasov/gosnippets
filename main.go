package main

import (
	"fmt"
	g "github.com/lanastasov/gosnippets/lib"
)

func main() {
	fmt.Println("----")
	// g.ScreenCls()
	g.Defer()
	fmt.Println(g.RepeatString("*", 8))
}
