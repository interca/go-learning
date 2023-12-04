package main

import "fmt"

func a() func() int {
	i := 0
	fmt.Println(&i)
	b := func() int {
		fmt.Println(&i)
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func main() {
	c := a()
	c()
	c()
	c()
	a() //不会输出i
}
