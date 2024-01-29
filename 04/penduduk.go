package main

import "fmt"

func main() {
	var population, births, immigration, deaths, emigration int

	fmt.Print("Masukkan jumlah penduduk awal: ")
	fmt.Scan(&population)
	fmt.Print("Masukkan jumlah kelahiran: ")
	fmt.Scan(&births)
	fmt.Print("Masukkan jumlah imigrasi: ")
	fmt.Scan(&immigration)
	fmt.Print("Masukkan jumlah kematian: ")
	fmt.Scan(&deaths)
	fmt.Print("Masukkan jumlah emigrasi: ")
	fmt.Scan(&emigration)

	finalPopulation := population + births + immigration - deaths - emigration

	fmt.Printf("Jumlah penduduk akhir: %d\n", finalPopulation)
}
