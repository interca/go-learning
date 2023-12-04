package main

import "fmt"

func c(x, y int) (z int) {
	defer func() {
		println(z) // 输出: 203
	}()

	z = x + y
	return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (return)
}

func main() {
	println(c(1, 2) + 1) // 输出: 203
	fmt.Println("---------------------------")
	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello, World!" },
	}
	fmt.Println(d.fn())

	a := struct {
		name string
	}{
		name: "hyj",
	}
	fmt.Println(a)
}
