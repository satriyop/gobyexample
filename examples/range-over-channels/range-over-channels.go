// Dalam contoh [sebelumnya](range) kita menggunakan
// `for` dan `range` untuk melakukan iterasi terhadap
// struktur data yang sederhana.
// Kita juga bisa menggunakan sintaks ini untuk melakukan
// iterasi terhadap value-value yang diterima dari sebuah
// channel.
package main

import "fmt"

func main() {

	// Kita akan melakukan iterasi terhadap 2 value
	// pada channel `queue`.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// `Range` akan melakukan iterasi terhadap setiap
	// elemen ketika value diterima dari channel `queue`.
	// Karena kita melakukan `close` terhadap channel pada
	// kode di atas, maka iterasi akan berakhir setelah
	// menerima 2 elemen saja.
	for elem := range queue {
		fmt.Println(elem)
	}
}
