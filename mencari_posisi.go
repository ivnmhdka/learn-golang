package main

import "fmt"

func main() {
	var d0 int
	var v int
	var t int
	var hasil int

	fmt.Scan(&d0, &v, &t)
	hasil = (v * t) + d0
	fmt.Println(hasil)
}
