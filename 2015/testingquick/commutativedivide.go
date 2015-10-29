package main

import (
	"fmt"
	"testing/quick"
)

func main() {
	// START OMIT
	f := func(a, b int) int { return a / b }
	g := func(a, b int) int { return b / a }
	fmt.Println("Counter examples against commutativity:", quick.CheckEqual(f, g, nil))
	// END OMIT
}
