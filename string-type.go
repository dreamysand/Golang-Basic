package main

import "fmt"

func main() {
	// String itu adalah kumpulan karakter
	var (
		stringval1 string = "Hello"
		stringval2 string = "World"
		stringval3 string = stringval1 + stringval2
	)

	fmt.Printf("Teks nya adalah %v, tipe datanya %T \n", stringval1, stringval1)
	fmt.Printf("Teks nya adalah %v, tipe datanya %T \n", stringval2, stringval2)
	fmt.Printf("Teks nya adalah %v, tipe datanya %T \n", stringval3, stringval3)
}