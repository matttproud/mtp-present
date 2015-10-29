package main

import (
	"fmt"
	"log"
	"sort"
	"testing/quick"

	"github.com/ryszard/goskiplist/skiplist"
)

func skipListTest(in []int) bool {
	// Arrange
	var (
		found     int
		no        = len(in)
		reference = make([]int, no)
		skiplist  = skiplist.NewIntSet()
	)
	copy(reference, in)
	sort.Ints(reference)

	// Act
	for _, v := range in {
		skiplist.Add(v)
	}

	// Assert
	for it := skiplist.Iterator(); it.Next(); {
		got := it.Key().(int)
		// first invariant
		if found > no {
			log.Print("skiplist contains more items than were given as input")
			return false
		}
		// second invariant
		if want := reference[found]; got != want {
			log.Printf("skiplist at %d got %d, want %d", found, got, want)
			return false
		}
		found++
	}
	// third invariant
	if found < no {
		log.Printf("skiplist had insufficient elements: got %d, want %d", found, no)
		return false
	}

	return true
}

func main() {
	// START OMIT
	err := quick.Check(skipListTest, nil /* configuration */)
	fmt.Printf("Skip list invariant counter examples: %v", err)
	// END OMIT
}
