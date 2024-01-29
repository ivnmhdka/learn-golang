package main

import "fmt"

func main() {
	var angka1, angka2 int
	var hasil int
	var hasil_asli bool

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&angka1)

	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&angka2)

	hasil = angka2 % angka1

	hasil_asli = hasil == 0

	fmt.Println(hasil_asli)
}
