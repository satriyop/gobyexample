// Channel adalah sebuah pipa/saluran yang menghubungkan
// goroutine yang berjalan secara concurrent.
// Kita bisa mengirimkan value-value ke dalam channel
// dari satu goroutine dan menerima value-value tersebut
// pada goroutine yang lainnya.
package main

import "fmt"

func main() {

	// Kita membuat channel baru dengan `make(chan val-type)`.
	// Jenis type dari channel ini sesuai dengan value yang
	// akan dimasukkan ke dalamnya.
	messages := make(chan string)

	// Mengirimkan sebuah value pada sebuah channel dengan
	// sintaks `channel <- `. Di sini kita mengirimkan
	// "ping" pada channel `messages` dari sebuah goroutine
	// yang baru.
	go func() { messages <- "ping" }()

	// sintaks `<- channel` akan menerima value dari
	// sebuah channel. Di sini kita akan menerima value
	// "ping" yang kita kirimkan dari kode di atas dan
	// menampilkannya pada layar.
	msg := <-messages
	fmt.Println(msg)
}
