package main

import (
	"fmt"
)

func main() {
	var a []int
	ints := append(a, 1)
	fmt.Println(a)
	fmt.Println(ints)
	var b = make([]int, 3, 10)
	fmt.Println(len(b), " ", cap(b))
	fmt.Println(b[0:1])
	var array [4]int
	sl := array[:2:3]
	fmt.Println(len(sl), cap(sl))
	fmt.Println(array)
	sl2 := append(sl, 100, 100)
	fmt.Println(array)
	fmt.Printf("%p\n", &sl2[0])
	fmt.Printf("%p\n", &sl[0])
}
