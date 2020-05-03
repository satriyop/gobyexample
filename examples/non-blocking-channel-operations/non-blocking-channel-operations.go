// Proses mengirim dan menerima pada channel pada dasarnya
// bersifat blocking, tapi kita bisa menggunakan `select`
// dengan klausa `default` untuk mengimplementasikan
// pengiriman dan penerimaan yang _non-blocking_  bahkan
// beberapa klausa `select` yang non-blocking.
package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Di contoh ini adalah penerimaan non-blocking.
	// Bila value tersedia pada `messages` maka `select`
	// akan memilih klausa `<-messages`. Jika tidak maka
	// akan segera memilih klausa `default`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Pengiriman non-blocking juga hampir sama. Di sini
	// `msg` tidak bisa dikirimkan ke channel `messages`
	// karena channel tidak punya _buffer_ dan tidak mempunyai
	// penerima, sehingga klausa `default` yang akan dipilih.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// Kita bisa menggunakan beberapa klausa sebelum
	// klausa `default` untuk mengimplementasikan
	// beberapa select non-blocking. Di contoh ini
	// kita berusaha menerima secara non-blocking pada
	// kedua channel `messages` dan `signals`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
