package main

import (
	"fmt"
	"sort"
	"testing/quick"
)

func SliceCopy(data sort.IntSlice) sort.IntSlice {
	d := make(sort.IntSlice, len(data))
	copy(d, data)
	return d
}
func main() {
	// START OMIT
	BubbleSort := func(data sort.Interface) {
		for {
			var swapped bool
			for i := 1; i < data.Len(); i++ {
				if data.Less(i, i-1) {
					data.Swap(i-1, i)
					swapped = true
				}
			}
			if !swapped {
				break
			}
		}
	}
	f := func(data sort.IntSlice) sort.IntSlice { d := SliceCopy(data); sort.Sort(d); return d }
	g := func(data sort.IntSlice) sort.IntSlice { d := SliceCopy(data); BubbleSort(d); return d }
	fmt.Println("Counter examples against sort:", quick.CheckEqual(f, g, nil))
	// END OMIT
}
