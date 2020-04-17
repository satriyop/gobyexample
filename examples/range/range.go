//_range_ akan mengiterasi elemen-elemen dalam
// berbagai struktur data. Mari kita lihat penggunaannya
// dengan struktur data yang sudah kita pelajari.
package main

import "fmt"

func main() {

	// Di contoh ini kita menggunakan `range` untuk
	// menjumlahkan angka dalam sebuah _slice_.
	// teknik ini bisa juga digunakan untuk _array_.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// `range` pada _array_ dan _slice_ menyediakan
	// index dan value untuk setiap elemen. Contoh di atas
	// kita tidak membutuhkan index, sehingga kita mengabaikannya
	// dengan _blank identifier_ yaitu `_`. Tapi kadang kita membutuhkan
	// index.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// `range` pada map akan mengiterasi pasangan key/value.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` bisa digunakan juga untuk hanya mengiterasi key saja
	// pada sebuah map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range` pada string akan mengiterasi _Unicode code point_.
	// Nilai pertama pada index byte dari `rune` dan nilai kedua
	// adalah `rune` nya itu sendiri.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
