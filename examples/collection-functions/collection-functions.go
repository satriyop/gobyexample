// Sering, dalam program kita dimana kita ingin
// melakukan sebuah operasi pada koleksi data, seperti
// memilih semua item yang sesuai dengan predikat atau
// memetakan semua item pada sebuah koleksi data baru
// dengan sebuah fungsi custom.

// Pada beberapa bahasa pemrograman lain hal ini
// bisa dilakukan dengan menggunakan struktur data
//  dan algoritma [generic](http://en.wikipedia.org/wiki/Generic_programming).
// Go tidak mempunyai generics; adalah hal umum di Go
// untuk menyediakan collection function jika dan hanya
// jika dibutuhkan oleh program dan tipe data.

// Di sini kita membuat contoh collection function
// untuk slice dari `strings`. Kita bisa menggunakan
// contoh ini untuk membuat fungsi kita sendiri.
// Perhatikan bahwa pada beberapa kasus, akan lebih
// jelas dengan hanya melakukan manipulasi koleksi data secara
// inline daripada membuat dan memanggil fungsi helper.

package main

import (
	"fmt"
	"strings"
)

// Index akan mengembalikan nilai index pertama  yang cocok dari
// target string `t` , atau -1 bila tidak ada yang cocok.
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include akan mengembalikan nilai `true` bila target string t
// ada dalam slice.
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any akan mengembalikan nilai `true` bila salah satu string
// dalam slice cocok dengan predikat `f`.
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All akan mengembalikan nilai `true` bila semua string
// dalam slice cocok dengan predicate `f`.
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter akan mengembalikan sebuah nilai baru yang mengandung
// semua nilai string dalam slie yang cocok dengan predikat `f`.
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map akan mengembalikan sebuah nilai slice yang mengandung
// hasil dari mengaplikasikan fungsi `f` pada setiap string
// dalam slice awal.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	// Di sini kita akan mencoba berbagai collection function.
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	// Contoh-contoh di atas semuanya menggunakan fungsi
	// anonim, tapi kita juga bisa menggunakan fungsi bernama
	// dari tipe yang tepat.
	fmt.Println(Map(strs, strings.ToUpper))

}
