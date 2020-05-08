// Pada contoh sebelumnya kita telah melihat bagaimana
// mengatur state counter sederhana menggunakan [atomic operations](atomic-counters).
// Untuk state yang lebih rumit kita bisa menggunakan
// <em>[mutex](http://en.wikipedia.org/wiki/Mutual_exclusion)</em>
// untuk mengakses data secara aman dari beberapa
// goroutine yang berjalan.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	// Dalam contoh ini, `state` adalah sebuah map.
	var state = make(map[int]int)

	// `mutex` akan melakukan sinkronisasi akses
	// terhadap `state`.
	var mutex = &sync.Mutex{}

	// Kita akan merekam berapa banyak operasi
	// read dan write yang dilakukan.
	var readOps uint64
	var writeOps uint64

	// Kita akan menjalankan 100 goroutine untuk menjalankan
	// read pada state secara berulang, satu kali per mili detik
	// pada setiap goroutine.
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				// Untuk setiap read kita akan memilih
				// sebuah key untuk mengakses state,
				// `Lock()` pada `mutex` untuk memastikan
				// akses eksklusif pada `state`, read dari
				// value dengan key yang sesuai, lakukan `Unlock()`
				// dari mutex dan menaikkan nilai dari `readOps`.
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				// Menunggu satu mili detik antara operasi read.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Kita akan menjalankan 10 goroutine untuk mensimulasi
	// operasi write, menggunakan pola yang sama seperti
	// sebelumnya pada operasi read.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Memberikan waktu untuk 10 goroutine bekerja pada
	// `state` dan `mutex` selama satu detik.
	time.Sleep(time.Second)

	// Mengambil jumlah operasi dan melaporkan.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// Dengan sebuah lock `state`, akan menunjukkan
	// hasil akhir operasi ini saat telah selesai.
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
