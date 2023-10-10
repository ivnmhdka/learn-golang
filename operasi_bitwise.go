package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int64
	var b int64
	var c_hasil int64

	a = 9
	b = 5
	// operasi and (&)
	c_hasil = a & b

	fmt.Println("=====AND=====")
	fmt.Println("nilai a :", a, ", nilai binernya : ", strconv.FormatInt(a, 2))
	fmt.Println("nilai b :", b, ", nilai binernya : ", strconv.FormatInt(b, 2))
	fmt.Println("------------&")
	fmt.Println("nilai hasil : ", c_hasil, "nilai binernya : ", strconv.FormatInt(c_hasil, 2))

	// operasi or (|)
	c_hasil = a | b

	fmt.Println("=====OR=====")
	fmt.Println("nilai a :", a, ", nilai binernya : ", strconv.FormatInt(a, 2))
	fmt.Println("nilai b :", b, ", nilai binernya : ", strconv.FormatInt(b, 2))
	fmt.Println("------------&")
	fmt.Println("nilai hasil : ", c_hasil, "nilai binernya : ", strconv.FormatInt(c_hasil, 2))

	// operasi xor (^)
	c_hasil = a ^ b

	fmt.Println("=====OR=====")
	fmt.Println("nilai a :", a, ", nilai binernya : ", strconv.FormatInt(a, 2))
	fmt.Println("nilai b :", b, ", nilai binernya : ", strconv.FormatInt(b, 2))
	fmt.Println("------------&")
	fmt.Println("nilai hasil : ", c_hasil, "nilai binernya : ", strconv.FormatInt(c_hasil, 2))
}
