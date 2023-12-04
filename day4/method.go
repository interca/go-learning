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
func main() {
	fmt.Println(count(111))
	var p Person
	println(p.sum())
	p.change()
	fmt.Println(p)
}
