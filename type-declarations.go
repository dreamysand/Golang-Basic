package main

import "fmt"

func main() {
	 // Type declarations adalah membuat tipe data baru berdasarkan tipe data yang sudah ada, ya hampir sama lah kayak rune
	 // Cara buatnya:
	 // type new_data_type data_type
	 type nisn int16
	 type kelas string
	 // Untuk pemakaiannya dalam variabel atau const sama seperti tipe data biasa
	 var (
	 	kelas_udin kelas = "XI RPL 1"
	 	kelas_asep kelas = "XI RPL 2"
	 )
	 const (
	 	nisn_udin nisn = 755
	 	nisn_asep nisn = 907
	 )

	 fmt.Printf("Kelas : %v, tipe datanya %T, nisn : %d, tipe datanya %T \n", kelas_udin, kelas_udin, nisn_udin, nisn_udin)
	 fmt.Printf("Kelas : %v, tipe datanya %T, nisn : %d, tipe datanya %T \n", kelas_asep, kelas_asep, nisn_asep, nisn_asep)
}