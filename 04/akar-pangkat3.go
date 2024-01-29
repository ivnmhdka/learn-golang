package main

import "fmt"

func main() {
	var angka1, angka2 int
	var hasil_asli bool

	fmt.Print("Masukkan bilangan x: ")
	fmt.Scan(&angka1)

	fmt.Print("Masukkan bilangan y: ")
	fmt.Scan(&angka2)

	hasil_asli = angka2 == (angka1 * angka1 * angka1)
	fmt.Println(hasil_asli)
}
