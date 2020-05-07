// Mekanisme utama untuk mengatur _state_ di Go adalah
// komunikasi melalui channel. Kita telah melihatnya di contoh
// [worker pools](worker-pools). Namun ada beberapa pilihan
// lain untuk mengatur state. Di contoh ini kita akan melihat
// penggunaan package `sync/atomic` untuk _atomic counters_
// yang diakses oleh beberapa goroutine.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Kita akan menggunakan sebuah unsigned integer
	// sebagai representasi counter (nilainya selalu positif).
	var ops uint64

	// Sebuah _WaitGroup_ akan digunakan untuk menunggu
	// semua goroutine selesai.
	var wg sync.WaitGroup

	// Kita akan menjalankan 50 goroutine yang akan
	// menaikkan counter sampai 1000 kali.
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				// Untuk menaikkan nilai cuonter secara
				// _atomic_ kita menggunakan `AddUint64`
				// dan memberikan alamat memory counter
				// `ops` dengan sintaks `&`.
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	// Menunggu sampai semua gorouting selesai.
	wg.Wait()

	// Sekarang sudah aman untuk mengakses `ops` karena
	// kita tahu tidak ada lagi goroutine yang menulisnya.
	// Membaca atomic secara aman sewaktu atomic diupdate juga
	// dimungkinkan menggunakan fungsi seperti `atomic.LoadUint64`.
	fmt.Println("ops:", ops)
}
