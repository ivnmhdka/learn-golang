package main

import "fmt"

func main() {
	var input int

	fmt.Print("Masukkan bilangan dua digit : ")
	fmt.Scanln(&input)

	if input < 10 || input > 99 {
		fmt.Println("Masukkan bilangan dua digit")
		return
	}

	result := (input%10)*1000 + (input%10)*100 + (input/10)*10 + (input / 10)

	fmt.Println("Penggandaan digit : ", result)
}
