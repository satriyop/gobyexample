// Go mendukung _multiple return values_ (lebih dari satu value yang
// bisa dikembalikan dari pemanggilan sebuah fungsi).
// Fitur ini sering digunakan dalam Go secara _idiomatic_,
// sebagai contoh, untuk mengembalikan (return) hasil maupun
// error dari sebuah fungsi.
package main

import "fmt"

// Dalam _signature_ fungsi ini, `(int, int)` menunjukkan
// bahwa fungsi ini mengembalikan sejumlah 2 `int`.
func vals() (int, int) {
	return 3, 7
}

func main() {

	// Di contoh ini kita menyimpan 2 value yang berbeda
	// dari pemanggilan sebuah fungsi dengan menggunakan
	// _multiple assignment_.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Jika kita tidak mau menyimpan semua value yang dikembalikan
	// dari sebuah pemanggilan fungsi
	// kita bisa menggunakan _blank identifier_ `_`
	// untuk mengacuhkan value lainnya.
	_, c := vals()
	fmt.Println(c)
}
