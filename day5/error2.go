package main

import "fmt"

func test() {
	defer func() {
		fmt.Println(recover()) //有效
	}()
	defer recover()              //无效！
	defer fmt.Println(recover()) //无效！
	defer func() {
		func() {
			println("defer inner")
			recover() //无效！
		}()
	}()

	func(x int) {
		fmt.Println(1)
	}(1)
	panic("test panic")
}

func main() {
	test()
}
