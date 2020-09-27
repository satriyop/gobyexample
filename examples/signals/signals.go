// Kadang kita ingin program Go kita menangani [Unix signals](http://en.wikipedia.org/wiki/Unix_signal).
// sebagai contoh kita ingin sebuah server melakukan
// shutdown secara graceful ketika menerima sebuah `SIGTERM`, atau
// sebuah tool command-line supaya berhenti memproses input ketika
// menerima sebuah `SIGINT`. Di contoh ini kita akan menangani
// signal di Go dengan menggunakan channel.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Notifikasi sinyal di Go berjalan dengan mengirimkan
	// value `os.Signal` pada sebuah channel. Kita akan
	// membuat sebuah channel untuk menerima notifikasi
	// (kita juga akan membuat satu lagi channel untuk
	// notifikasi ketika program bisa berakhir)
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// `signal.Notify` me-register channel yang diberikan
	// untuk menerima notifikasi dari sinyal yang telah
	// ditentukan.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine ini akan mengeksekusi sebuah receive
	// blocking untuk sinyal-sinyal. Ketika menerima
	// sebuah sinyal akan menampilkan dan menotifikasi
	// program bahwa program bisa berakhir.
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// Program akan menunggu di sini sampai
	// menerima sinyal yang diharapkan (sebagaimana
	// diindikasikan oleh goroutine di atas mengirimkan value
	// pada channel `done`) dan kemudian mengakhiri program.
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
