package main

import "fmt"

func main() {
	var mass_kg float64
	var rumus float64
	var merkurius, venus, bumi, mars, yupiter, saturnus, uranus, neptunus float64

	fmt.Print("Masukkan berat: ")
	fmt.Scan(&mass_kg)
	rumus = mass_kg * 9.8

	merkurius = rumus * 0.38
	venus = rumus * 0.91
	bumi = rumus
	mars = rumus * 0.39
	yupiter = rumus * 2.37
	saturnus = rumus * 0.92
	uranus = rumus * 0.89
	neptunus = rumus * 1.13

	fmt.Println("Berat tubuh manusia di berbagai planet:")
	fmt.Printf("Merkurius: \n", merkurius)
	fmt.Printf("Venus: \n", venus)
	fmt.Printf("Bumi: \n", bumi)
	fmt.Printf("Mars: \n", mars)
	fmt.Printf("Yupiter: \n", yupiter)
	fmt.Printf("Saturnus: \n", saturnus)
	fmt.Printf("Uranus: \n", uranus)
	fmt.Printf("Neptunus: \n", neptunus)
}
