package main

import "fmt"

func main() {
	var nilai_mhs int
	var tugas bool
	var status bool

	fmt.Scan(&nilai_mhs, &tugas) // input
	status = (nilai_mhs > 55 && tugas == true) || nilai_mhs > 90
	fmt.Println(status) // output
}
