// Go mendukung
// <a href="http://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>recursive functions</em></a>.
// Berikut adalah contoh klasik dari perhitungan faktorial.
package main

import "fmt"

// Contoh fungsi `fact` ini memanggil dirinya sendiri
// sampai mencapai _base case_ yaitu `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
