package main

import "fmt"

func count(a int) int {
	return a * 10
}

type Person struct {
	name string
}

func (p Person) sum() int {
	return 1
}

func (p *Person) change() {
	p.name = "黄"
}

func add(a ...int) []int {
	a[0] = 10
	return a
}

func main() {
	fmt.Println(count(111))
	var p Person
	println(p.sum())
	p.change()
	fmt.Println(p)
	ints := add(1, 2, 3)
	fmt.Println(ints)
}
