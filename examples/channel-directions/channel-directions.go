// Ketika kita menggunakan channel sebagai parameter
// dari sebuah fungsi, maka kita bisa menentukan jika
// sebuah channel hanya untuk mengirim atau hanya untuk
// menerima value. Dengan melakukannya, kita meningnkatkan
// _type-safety_ dari program kita.
package main

import "fmt"

// Fungsi `ping` ini hanya akan menerima channel untuk
// mengirimkan value. Aka terjadi error pada _compile-time_
// jika kita mencoba untuk mengirim value pada channel ini.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Parameter pertama fungsi `pong` ini adalah sebuah channel untuk
// menerima (`pings`) dan parameter kedua untuk mengirim (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
