package main

import "fmt"

func main() {
	var x float64

	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&x)

	fraction := ((x*x*x + 3*x) / (x*x*x*x - 3*x*x + 4))
	fmt.Print("Nilai dari f(x) : ", fraction)
}
