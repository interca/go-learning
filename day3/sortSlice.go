package main

import (
	"fmt"
	"sort"
)

func main() {
	var num = []int{4, 3, 21}
	var s = num[:]
	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)
	fmt.Println(num)
}
