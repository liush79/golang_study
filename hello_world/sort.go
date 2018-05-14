package main

import (
	"fmt"
	"sort"
)

func testSort() {
	a := []int{10, 5, 3, 7, 6}
	sort.Ints(a)
	fmt.Println(a)

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
}
