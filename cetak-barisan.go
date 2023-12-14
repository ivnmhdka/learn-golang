package main

import "fmt"

func main() {
	var i int
	var a, b int
	a = 1
	b = 10

	for i = a; i <= b; i++ {
		fmt.Println(i * 2)
	}
}
