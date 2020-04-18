// _Function_ adalah pusat dari Go. Kita akan mempelajari
// function dengan contoh-contoh yang berbeda.
package main

import "fmt"

// Di contoh ini, sebuah function akan mengambil dua
// parameter dengan type `int`
// dan mengembalikan hasil penjumlahan keduanya
// dengan type `int` juga.
func plus(a int, b int) int {

	// Go membutuhkan _keyword_ return secara eksplisit,
	// artinya, Go tidak mengembalikan nilai
	// dari _expression_ terakhir secara otomatis.
	return a + b
}

// Ketika kita punya beberapa parameter dengan type yang sama,
// kita bisa mendeklarasikannya secara sekaligus.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// Memanggil function seperti yang kita bayangkan,
	// dengan `nama(args)`.
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
