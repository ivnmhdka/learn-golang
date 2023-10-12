package main

import (
	"fmt"
)

func main() {
	var x int
	var d1, d2, d3 int

	fmt.Print("Masukkan bilangan bulat positif (kurang dari atau sama dengan 999) : ")
	fmt.Scan(&x)

	d1 = (x / 100)
	fmt.Printf("Nilai dari d1 : %d", d1)

	d2 = (x / 10) % 10
	fmt.Printf(" Nilai dari d2 : %d", d2)

	d3 = (x / 1) % 10
	fmt.Printf(" Nilai dari d3 : %d", d3)
}
