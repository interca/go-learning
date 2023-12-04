package main

import "fmt"

// 外部引用函数参数局部变量
func add2(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func main() {
	tmp1 := add2(10)
	fmt.Println(tmp1(1), tmp1(2))
	// 此时tmp1和tmp2不是一个实体了
	tmp2 := add2(100)
	fmt.Println(tmp2(1), tmp2(2))
}
