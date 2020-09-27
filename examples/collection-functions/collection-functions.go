// We often need our programs to perform operations on
// collections of data, like selecting all items that
// satisfy a given predicate or mapping all items to a new
// collection with a custom function.

// Sering, dalam program kita dimana kita ingin
// melakukan sebuah operasi pada koleksi data, seperti
// memilih semua item yang sesuai dengan predikat atau
// memetakan semua item pada sebuah koleksi data baru
// dengan sebuah fungsi custom.

// In some languages it's idiomatic to use [generic](http://en.wikipedia.org/wiki/Generic_programming)
// data structures and algorithms. Go does not support
// generics; in Go it's common to provide collection
// functions if and when they are specifically needed for
// your program and data types.

// Pada beberapa bahasa pemrograman lain hal ini
// bisa dilakukan dengan menggunakan struktur data
//  dan algoritma [generic](http://en.wikipedia.org/wiki/Generic_programming).
// Go tidak mempunyai generics; adalah hal umum di Go
// untuk menyediakan collection function jika dan hanya
// jika dibutuhkan oleh program dan tipe data.

// Here are some example collection functions for slices
// of `strings`. You can use these examples to build your
// own functions. Note that in some cases it may be
// clearest to just inline the collection-manipulating
// code directly, instead of creating and calling a
// helper function.

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

// Index returns the first index of the target string `t`, or
// -1 if no match is found.

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

// Include returns `true` if the target string t is in the
// slice.

// Include akan mengembalikan nilai `true` bila target string t
// ada dalam slice.
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any returns `true` if one of the strings in the slice
// satisfies the predicate `f`.

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

// All returns `true` if all of the strings in the slice
// satisfy the predicate `f`.

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

// Filter returns a new slice containing all strings in the
// slice that satisfy the predicate `f`.

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

// Map returns a new slice containing the results of applying
// the function `f` to each string in the original slice.

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

	// Here we try out our various collection functions.

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

	// The above examples all used anonymous functions,
	// but you can also use named functions of the correct
	// type.

	// Contoh-contoh di atas semuanya menggunakan fungsi
	// anonim, tapi kita juga bisa menggunakan fungsi bernama
	// dari tipe yang tepat.
	fmt.Println(Map(strs, strings.ToUpper))

}
