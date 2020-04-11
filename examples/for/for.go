// `for` adalah satu-satunya _construct_ looping di Go.
// Berikut adalah tiga tipe perulangan dasar `for`
package main

import "fmt"

func main() {

	// Tipe perulangan paling dasar, menggunakan satu persyaratan.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Tipe perulangan `for` klasik dengan format
	// inisiasi/persyaratan/langkah
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` tanpa persyaratan akan terus berulang
	// sampai ada `break` dari perulangan itu atau
	// `return` dari function yang menaungi
	for {
		fmt.Println("loop")
		break
	}

	// Kita bisa menggunakan `continue` untuk melanjutkan
	// ke iterasi berikutnya.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
