package main

import "fmt"

func main() {
	var suku1, suku2 int

	fmt.Print("Masukkan suku ke-1: ")
	fmt.Scan(&suku1)
	fmt.Print("Masukkan suku ke-2: ")
	fmt.Scan(&suku2)

	suku3 := suku1 + suku2
	suku4 := suku2 + suku3
	suku5 := suku3 + suku4

	fmt.Printf("Suku ke-3: %d\n", suku3)
	fmt.Printf("Suku ke-4: %d\n", suku4)
	fmt.Printf("Suku ke-5: %d\n", suku5)
}
