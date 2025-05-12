package main

import "fmt"

func main() {
	// Untuk operasi matematika sebenarnya sama saja seperti kebanyakan bahasa pemrograman
	// Ada:
	// 1. Penjumlahan
	a := 10
	b := 10
	c := a+b
	fmt.Printf("Hasil dari %d + %d = %d \n", a, b, c)
	// 2. Pengurangan
	a = 10
	b = 10
	c = a-b
	fmt.Printf("Hasil dari %d - %d = %d \n", a, b, c)
	// 3. Perkalian
	a = 10
	b = 10
	c = a*b
	fmt.Printf("Hasil dari %d * %d = %d \n", a, b, c)
	// 4. Pembagian
	a = 10
	b = 10
	c = a/b
	fmt.Printf("Hasil dari %d / %d = %d \n", a, b, c)
	// 5. Modulus => Hasil bagi
	a = 10
	b = 10
	c = a%b
	fmt.Printf("Hasil dari %d mod %d = %d \n", a, b, c)

	println()
	/*
	Augmented operators
	 */
	x := 5
	x += 10 //x + 10
	fmt.Printf("Hasil dari %d + 10 = %d \n", x, x)
	x -= 10 //x - 10
	fmt.Printf("Hasil dari %d - 10 = %d \n", x, x)
	x *= 10 //x * 10
	fmt.Printf("Hasil dari %d * 10 = %d \n", x, x)
	x /= 10 //x / 10
	fmt.Printf("Hasil dari %d / 10 = %d \n", x, x)
	x %= 10 //x % 10
	fmt.Printf("Hasil dari %d mod 10 = %d \n", x, x)
}