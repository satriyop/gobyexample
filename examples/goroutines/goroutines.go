// Goroutine adalah sebuah eksekusi thread yang ringan.
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Ketika kita mempunyai fungsi dengan nama `f(s)`,
	// Kita memanggilnya fungsi tersebut seperti ini,
	// sehingga fungsi ini berjalan secara _synchronous_.
	f("direct")

	// Untuk memanggil fungsi ini dalam sebuah
	// goroutine, gunakan `go f(s)`. Goroutine yang
	// baru ini akan mengeksekusi fungsi ini secara
	// `concurrent`.
	go f("goroutine")

	// Kita juga bisa menjalankan sebuah goroutine pada
	// sebuah fungsi _anonymous_ (tanpa nama) seperti ini.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Kedua fungsi ini akan berjalan secara `asynchronous`
	// dalam goroutine yang berbeda (terpisah).
	// Kita akan menunggu goroutine-goroutine ini untuk selesai
	// (kalau tidak fungsi main akan langsung menutup aplikasi kita).
	// Untuk methode menunggu goroutine yang lebih bagus, gunakan
	// [WaitGroup](waitgroups).
	time.Sleep(time.Second)
	fmt.Println("done")
}
