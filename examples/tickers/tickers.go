//[Timers](timers)  dipakai untuk sesuatu yang ingin kita
// lakukan saat nanti, sedangkan _ticker_ digunakan ketika
// kita ingin melakukan sesuatu secara berulang pada sebuah
// interval. Di contoh ini adalah ticker yang berjalan
// secara periodis sampai kita menghentikannya.
package main

import (
	"fmt"
	"time"
)

func main() {

	// Ticker menggunakan mekanisme serupa dengan timer:
	// sebuah channel yang mengirimkan value. Di sini kita
	// menggunakan `select` pada channel untuk menunggu
	// value diterima setiap 500 mili detik.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Ticker bisa dihentikan seperti timer. Ketika ticker
	// berhenti maka tidak akan menerima lagi value pada
	// channel. Kita akan menghentikan ticker kita setelah
	// 1600 mili detik.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
