package main

// contoh program penjumlahan

import "fmt"

func main() {
	var a int
	var b int
	var hasil int     // ini tipe data integer
	var kabar string  // ini tipe data string
	var nilai float32 // ini tipe data real
	var setuju bool   // ini tipe data boolean

	kabar = "kabar baik"
	nilai = 3.0
	setuju = true

	a = 16
	b = 20
	hasil = a + b
	fmt.Println("hasil penjumlahan = ", hasil)
	fmt.Println(kabar)
	fmt.Println("nilai tipe data real = ", nilai)
	fmt.Println("nilai tipe data boolean = ", setuju)

}
