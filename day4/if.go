package main

import "fmt"

func main() {
	var a int
	if a > 1 && a != 2 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	var num = []int{1, 2, 3, 4}
	var k, v int
	for k, v = range num {
		fmt.Println(k, v)
	}
}
