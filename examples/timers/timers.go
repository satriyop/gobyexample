// Kita sering ingin mengeksekusi sebuah kode di Go
// di saat nanti, atau berulang pada sebuah interval.
// _timer_ dan _ticker_ adalah fitur dari Go yang
// membuat kedua hal tersebut mudah. Kita akan melihat
// timer terlebih dahulu kemudian [tickers](tickers).
package main

import (
	"fmt"
	"time"
)

func main() {

	// Timer merepresentasikan sebuah event di masa datang.
	// Kita bisa mengatur seberapa lama timer harus menunggu,
	// dan timer menyediakan sebuah channel yang akan diberitahu
	// bila waktunya sudah habis. Timer ini akan menunggu selama
	// 2 detik saja.
	timer1 := time.NewTimer(2 * time.Second)

	// kode `<-timer.C` membuat block pada channel `C` pada timer
	// sampai mengirimkan sebuah value yang mengindikasikan bahwa
	// timer telah dimulai.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// Jika kita ingin menunggu, kita bisa menggunakan
	// `time.Sleep`. Satu alasan kenapa sebuah timer berguna
	// adalah kita bisa membatalkan sebuah timer sebelum
	// timer tersebut dimulai. Berikut ini contohnya.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Memberikan `timer2` cukup waktu untuk mulai, bilapun
	// akan terjadi, tapi di program ini akan menunjukkan
	// bahwa timer telah berhenti.
	time.Sleep(2 * time.Second)
}
