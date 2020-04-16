// _Maps_ adalah [associative data type](http://en.wikipedia.org/wiki/Associative_array) di Go
// (kadanng dinamakan _hash_ atau _dictionary_ di bahasa pemrograman lain).
package main

import "fmt"

func main() {

	// Untuk membuat map kosong, gunakan fungsi builtin `make`
	// `make(map[type-dari-key]type-dari-value)`
	m := make(map[string]int)

	// Menatur pasangan key/value dengan
	// sintaks `nama[key] = value`
	m["k1"] = 7
	m["k2"] = 13

	// Menampilkan map dengan `fmt.Println` akan menunjukkan
	// semua pasangan key/value
	fmt.Println("map:", m)

	// Mengambil sebuah value dari key dengan `nama[key]`
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// `len` dengan argumen map akan mengambalikan jumlah dari pasangan
	// key/value dari map tersebut.
	fmt.Println("len:", len(m))

	// `delete` dengan argumen pertama map, dan argumen kedua key
	// akan menghapus pasangan key/value tersebut dari map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// Return value kedua (yang opsional) ketika mengakses map
	// mengindikasikan bahwa key tersebut ada di dalam map.
	// Hal ini bisa digunakan untuk menghindari ambigu antara
	// key yang tidak ada dan key yang bernilai _zero values_
	// seperti `0` atau `""`. Di contoh ini, kita tidak membutuhkan
	// valuenya, sehingga menggunakan _blank identifier_
	// pada return value pertama dengan notasi `_`.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Kita juga bisa mendeklarasikan dan menginisiasikan
	// map baru dalam satu baris dengan sintaks berikut ini.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
