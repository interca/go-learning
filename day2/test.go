package main

import (
	"fmt"
	_ "go-learning/day1"
	"strings"
)

const b = 1

func main() {
	fmt.Println(1)
	var a = true
	fmt.Println(a)
	fmt.Println(b, a)
	var s = "ac34567"
	s2 := s[1:len(s)]
	fmt.Println(s2)
	w := 1
	fmt.Println(w)
	var c = s[0]
	fmt.Println(c)
	bytes := []byte(s)
	bytes[0] = 'c'
	s = string(bytes)
	fmt.Println(s)
	replace := strings.Replace(s, string(s[0]), "黄", 3)
	fmt.Println(replace)
}
