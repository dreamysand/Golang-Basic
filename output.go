package main

import "fmt"

func main() {
	var (
		hello = "hello"
		world = "world"
	)
	// Ada 3 macam cara output:
	// 1. Print => defaultnya, tanpa ditambahin apa apa intinya apa adanya
	// untuk melakukan print ada 2 cara:
	// - Import fmt
	fmt.Print(hello, world, "\n")
	// - Langsung print()
	print(hello, world, "\n")
	// 2. Println -> Juga print new line
	// untuk melakukan println juga ada 2 cara:
	// - Import fmt
	fmt.Println(hello, world)
	// - Langsung println
	println(hello, world)
	// 3. Printf => print dengan langsung menggunakan format variabelnya misal
	// fmt.Printf("halo %v", world) => %v ini format yang akan digantikan variabel world
	fmt.Printf("Halo %v \n", world)
	fmt.Printf("%v World \n", hello)
}