package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}

	for i, v := range s { // 复制 struct slice { pointer, len, cap }。
		fmt.Println(v)
		if i == 0 {
			s = s[:3]  // 对 slice 的修改，不会影响 range。
			s[2] = 100 // 对底层数据的修改。
		}
	}

out:
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if i+j == 3 {
				break out
			}
			fmt.Println(i + j)
		}
	}
}
