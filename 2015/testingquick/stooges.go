package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing/quick"
)

// START OMIT
type Stooge int

const (
	Invalid Stooge = iota
	Moe
	Larry
	Shemp
	Curly
	Joe
	CurlyJoe

	nStooges = int(CurlyJoe) + 1
)

func (s Stooge) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Stooge(rand.Intn(nStooges)))
}

// END OMIT

func (s Stooge) String() string {
	switch s {
	case Invalid:
		return "Invalid"
	case Moe:
		return "Moe"
	case Larry:
		return "Larry"
	case Shemp:
		return "Shemp"
	case Curly:
		return "Curly"
	case Joe:
		return "Joe"
	case CurlyJoe:
		return "Curly Joe"
	}
	panic("unreachable")
}

func main() {
	var (
		rnd = rand.New(rand.NewSource(42))
		t   = reflect.TypeOf(Stooge(0))
	)
	fmt.Println("Introducing the Stooges:")
	for i := 0; i < 3; i++ {
		v, _ := quick.Value(t, rnd)
		fmt.Println(v.Interface())
	}
}
