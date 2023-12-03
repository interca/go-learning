package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	name string
	age  int
}

func main() {
	var a person
	a.name = "hyj黄"
	fmt.Println(a)
	b := &person{}
	fmt.Println(b)
	fmt.Println("----------------------")

	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
