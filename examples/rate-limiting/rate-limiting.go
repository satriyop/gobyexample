// <em>[Rate limiting](http://en.wikipedia.org/wiki/Rate_limiting)</em>
// adalah mekanisme penting untuk mengontrol penggunaan
// resource dan menjaga kualitas layanan.
// Go secara elegan mendukung rate limiting dengan
// goroutine, channel, dan [tickers](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// Pertama mari kita lihat rate limiting sederhana.
	// Bila kita ingin membatasi penerimaan request yang datang.
	// Kita akan melayani request dari sebuah channel
	// dengan nama yang sama.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// channel `limiter` ini akan menerima value setiap
	// 200 mili detik. Hal ini adalah pengatur dalam
	// skema rate limiting kita.
	limiter := time.Tick(200 * time.Millisecond)

	// Dengan melakukan block pada setiap penerimaan dari channel
	// `limiter` sebelum melayani setiap request, kita membatasi
	// 1 request setiap 200 mili detik.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Kita mungkin ingin menerima lonjakan pendek request
	// dalam skema rate limiting dengan tetap mempertahankan
	// rate limiting secara umum. Kita bisa mencapai ini
	// dengan memberikan buffer pada channel limiter.
	// Channel `burstyLimiter` ini akan mengijinkan lonjakan
	// sebanyak 3 kali.
	burstyLimiter := make(chan time.Time, 3)

	// Mengisi channel untuk merepresentasikan lonjakan
	// yang diijinkan.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Setiap 200 mili detik kita akan mencoba
	// menambah value pada `burstyLimiter`, sampai
	// batasnya yaitu 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Sekarang kita mensimulasikan tambahan
	// 5 request datang. Ketiga pertamanya akan menerima
	// keuntungan dari `burstyLimiter`.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
