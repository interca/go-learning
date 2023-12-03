package main

import "fmt"

func main() {
	var m = make(map[string]string, 10)
	m["hyj"] = "123"
	for _, s := range m {
		fmt.Println(s)
	}
	delete(m, "111")
}
