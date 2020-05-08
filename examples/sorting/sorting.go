// Package `sort` di Go mengimplementasikan proses
// sorting (pengurutan)  pada type built-in maupun user-defined.
// Kita akan melihat bagaimana sorting untuk type built-in
// terlebih dahulu.
package main

import (
	"fmt"
	"sort"
)

func main() {

	// Method-method dari package sort adalah spesifik
	// untuk type builtin . Perhatikan bahwa sorting
	// melakukan operasinya dengan mengubah slice yang
	// yang diberikan secara langsung dan
	// tidak mengembalikan slice baru.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// Contoh dari mengurutkan type `int`.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// Kita juga bisa menggunakan `sort` untuk mengecek
	// bahwa sebuah slice telah urut.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
