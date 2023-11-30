package main

import "fmt"

func change(a *[4]int) {
	a[0] = 1
}

type node struct {
	a int
	b int
}

func main() {
	var a [4]int
	change(&a)
	var i, v int
	for i, v = range a {
		fmt.Println(i, " ", v)
	}
	var b = [...]int{1, 2, 3}
	fmt.Println(b[0])
	fmt.Println()
	var c = [...]node{
		{1, 2},
		{3, 4},
	}
	for s := range c {
		fmt.Println(c[s])
	}
	fmt.Println()
	var arr0 = [1][3]int{
		{1, 2, 3},
	}
	fmt.Println(arr0)
}
