package main

import "fmt"

func main() {
	var a []int
	ints := append(a, 1)
	fmt.Println(a)
	fmt.Println(ints)
	var b = make([]int, 3, 10)
	fmt.Println(len(b), " ", cap(b))
	fmt.Println(b[0:1])
}
