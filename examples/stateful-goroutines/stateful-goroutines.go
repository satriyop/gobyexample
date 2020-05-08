// Dalam contoh sebelumnya kita menggunakan _locking_
// secara eksplisit menggunakan [mutex](mutexes) untuk
// sinkronisasi beberapa goroutine yang mengakses state.
// Opsi lainnya adalah menggunakan fitur sinkronisasi
//  _built-in_ dari goroutine dan channel untuk mendapatkan
// hasil yang sama. Pendekatan dengan menggunakan channel
// ini sesuai dengan ide dari Go yaitu berbagi memory dengan
// komunikasi dan membuat satu buah data dimiliki secara
// persis oleh 1 goroutine.
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Dalam contoh ini state kita akan dimiliki oleh
// sebuah goroutine. Hal ini akan menjamin bahwa data
// tidak akan pernah rusak karena akses yang concurrent.
// Untuk melakukan read atau write pada state, goroutine lain
// akan mengirimkan message kepada goroutine yang memiliki
// state tersebut dan menerima jawaban. Struct `readOp` dan
// `writeOp` mengenkapsulasi request-request tersebut dan
// merupakan cara dari goroutine pemilik state merespon
// request.
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// Sebagaimana sebelumnya, kita akan menghitung jumlah
	// operasi yang terjadi.
	var readOps uint64
	var writeOps uint64

	// Channel `reads` dan `writes` akan digunakan goroutine
	// lain untuk membuat request read dan write.
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Ini adalah goroutine sebagai pemilik dari `state`,
	// yang berupa map seperti contoh sebelumnya tapi kali ini
	// state ini hanya private untuk stateful goroutine.
	// Goroutine ini akan secara berulang melakukan select
	// pada channel `reads` dan `writes` untuk melakukan
	// respon saat ada request. Sebuah respon dieksekusi
	// sesuai dengan operasi yang diminta yang terlebih dulu ada.
	// dan mengirimkan sebuah value pada channel respon `resp`
	// untuk mengindikasikan operasi berhasil (dan memberikan
	// value yang diminta bila operasinya adalah `reads`).
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Memulai 100 goroutine untuk membuat operasi read pada
	// goroutine pemilik state dengan menggunakan channel `reads`.
	// Setiap read membutuhkan proses kontstruksi dari struct
	// `readOp`, mengirimkannya pada channel `reads` dan menerima
	// hasilnya pada channel `resp`.
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Kita mulai operasi 10 writes dengan pendekatan
	// yang sama.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Membuat goroutine berjalan selama satu detik.
	time.Sleep(time.Second)

	// Akhirnya, mengambil dan melaporkan perhitungan
	// jumlah operasi
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
