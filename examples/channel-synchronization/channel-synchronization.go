// Kita bisa menggunakan channel untuk melakukan
// sinkronisasi dari goroutine-goroutine yang berjalan.
// Di sini kita menggunakan _blocking receive_ untuk
// menunggu dari sebuah goroutine selesai,
// Kita mungkin akan lebih memilih metode sinkronisasi
// dengan [WaitGroup](waitgroups) daripada metode ini.
package main

import (
	"fmt"
	"time"
)

// Ini adalah fungsi yang akan kita jalankan dalam
// sebuah goroutine. Channel `done` akan digunakan
// untuk memberitahu goroutine lain (func main juga adalah goroutine)
// bahwa fungsi ini telah selesai dieksekusi.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Mengirimkan sebuah value untuk memberitahu
	// bahwa fungsi ini sudah selesai.
	done <- true
}

func main() {

	// Memalai goroutine dari fungsi worker, dan memberikan
	// channel untuk pemberitahuan selesai nantinya.
	done := make(chan bool, 1)
	go worker(done)

	// Dengan melakukan ini, kita melakukan block sampai
	// ada pemberitahuan dari worker telah selesai.
	<-done
}
