// [_Variadic functions_](http://en.wikipedia.org/wiki/Variadic_function)
// bisa dipanggil dengan sejumlah argumen yang bervariasi.
// `fmt.Println` adalah salah satu contoh dari fungsi variadic.
package main

import "fmt"

// Ini adalah sebuah contoh fungsi variadic yang akan menerima sejumlah `int`
// sebagai argumennya (jumlah dari argumen tidak ditentukan).
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// Fungsi variadic bisa dipanggil seperti memanggil fungsi biasa
	// tapi dengan jumlah argumen yang bisa berbeda-beda.
	sum(1, 2)
	sum(1, 2, 3)

	// Jika kita sudah mempunyai sejumlah argumen dalam bentuk slice,
	// kita bisa mengaplikasikannya dalam sebuah pemanggilan fungsi variadic
	// dengan menggunakan sintaks `func(slice...)` .
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
