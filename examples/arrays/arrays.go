// Dalam bahasa pemrograman Go, sebuah _array_ adalah
// sebuah kumpulan elemen dengan nomor urutan
// yang mempunyai besaran yang sudah ditentukan.
package main

import "fmt"

func main() {
	// Kita membuat sebuah array `a` yang akan menampung
	// 5 buah `int`. Tipe dari elemen-element dan panjang
	// array adalah bagian dari type array. Secara default
	// sebuah array adalah _zero-valued_ dimana dalam hal ini
	// yaitu `int` berarti nilai-nilai elemennya adalah 0 semua.
	var a [5]int
	fmt.Println("emp:", a)

	// We can set a value at an index using the
	// `array[index] = value` syntax, and get a value with
	// `array[index]`.

	// Kita bisa mengatur sebuah nilai dari elemen dengan
	// menggunakan indeks atau nomoer urutannya dengan sintaks
	// `array[index] = value` , dan bisa mengakses nilai element
	// dengan `array[index]`
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// fungsi _builtin_ `len` akan me-return panjang dari sebuah array
	fmt.Println("len:", len(a))

	// Gunakan sintaks ini untuk mendeklarasikan dan menginisiasikan
	// sebuah array dalam satu baris.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Type array adalah satu dimensi, tapi kita bisa
	// mengusun kumpulan type untuk membuat data struktur
	// multi dimensi.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
