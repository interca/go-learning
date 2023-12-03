package main

import "fmt"

func main() {
	var a *int
	var b = 1
	a = &b
	fmt.Println(*a)

	var c map[string]int
	c = make(map[string]int, 1)
	c["测试"] = 100
	c["ss"] = 1
	fmt.Println(c)
}
