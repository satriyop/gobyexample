// Go mendukung [_anonymous functions_](http://en.wikipedia.org/wiki/Anonymous_function),
// yang bisa digunakan untuk membuat sebuah
// fungsi anonim berguna ketika kita ingin mendefinisikan
// sebuah fungsi secara _inline_ tanpa harus memberi nama fungsi tersebut.
package main

import "fmt"

// Fungsi `intSeq` ini mengembalikan suatu fungsi lain,
// yang mana kita definisikan secara anonim pada body fungsi `intSeq`.
// Fungsi anonim yang dikembalikan akan _menampung_ variabel `i`
// untuk membentuk sebuah closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// Kita memanggil fungsi `intSeq` dan menyimpan hasilnya
	// (yang adalah sebuah fungsi anonim) ke variabel `nextInt`.
	// Fungsi ini akan mempunyai value `i` sendiri, yang akan
	// berganti setiap kita memanggil fungsi `nextInt` ini.
	nextInt := intSeq()

	// Lihat efek dari _closure_ dengan memanggil fungsi
	// `nextInt` beberapa kali.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Untuk mengkonfirmasi bahwa _state_ adalah unik
	// terhadap fungsi tersebut, buatlah lagi fungsi baru
	// dan lakukan pengetesan seperti ini.
	newInts := intSeq()
	fmt.Println(newInts())
}
