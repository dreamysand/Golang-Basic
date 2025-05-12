package main

import "fmt"

func main() {
	// Operasi perbandingan => untuk membandingkan dua variabel => Kalau benar hasilnya true kalau salah false
	// Beberapa operator perbandingan:
	// 1. > : Lebih besar dari
	var (
		condition1 bool
		condition2 bool
		condition3 bool
	)
	condition1 = 10 > 5
	condition2 = 10 > 15
	condition3 = 10 > 10

	fmt.Printf("Pernyataan 10 > 5 nilainya adalah %v \n", condition1)
	fmt.Printf("Pernyataan 10 > 15 nilainya adalah %v \n", condition2)
	fmt.Printf("Pernyataan 10 > 10 nilainya adalah %v \n", condition3)
	println()
	// 2. < : Lebih kecil dari
	condition1 = 10 < 5
	condition2 = 10 < 15
	condition3 = 10 < 10

	fmt.Printf("Pernyataan 10 < 5 nilainya adalah %v \n", condition1)
	fmt.Printf("Pernyataan 10 < 15 nilainya adalah %v \n", condition2)
	fmt.Printf("Pernyataan 10 < 10 nilainya adalah %v \n", condition3)
	println()
	// 3. >= : Lebih besar sama dengan dari
	condition1 = 10 >= 5
	condition2 = 10 >= 15
	condition3 = 10 >= 10

	fmt.Printf("Pernyataan 10 >= 5 nilainya adalah %v \n", condition1)
	fmt.Printf("Pernyataan 10 >= 15 nilainya adalah %v \n", condition2)
	fmt.Printf("Pernyataan 10 >= 10 nilainya adalah %v \n", condition3)
	println()
	// 4. <= : Lebih kecil sama dengan dari
	condition1 = 10 <= 5
	condition2 = 10 <= 15
	condition3 = 10 <= 10

	fmt.Printf("Pernyataan 10 <= 5 nilainya adalah %v \n", condition1)
	fmt.Printf("Pernyataan 10 <= 15 nilainya adalah %v \n", condition2)
	fmt.Printf("Pernyataan 10 <= 10 nilainya adalah %v \n", condition3)
	println()
	// 5. == : Sama dengan
	condition1 = 10 == 5
	condition2 = 10 == 15
	condition3 = 10 == 10

	fmt.Printf("Pernyataan 10 == 5 nilainya adalah %v \n", condition1)
	fmt.Printf("Pernyataan 10 == 15 nilainya adalah %v \n", condition2)
	fmt.Printf("Pernyataan 10 == 10 nilainya adalah %v \n", condition3)
	println()
	// 6. != : Tidak sama dengan
	condition1 = 10 != 5
	condition2 = 10 != 15
	condition3 = 10 != 10

	fmt.Printf("Pernyataan 10 != 5 nilainya adalah %v \n", condition1)
	fmt.Printf("Pernyataan 10 != 15 nilainya adalah %v \n", condition2)
	fmt.Printf("Pernyataan 10 != 10 nilainya adalah %v \n", condition3)
	println()
}