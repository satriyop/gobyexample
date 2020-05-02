// Secara default channel tidak mempunyai buffer (_unbuffered_),
// artinya channel hanya akan menerima kiriman (`chan <-`) bila
// ada penerima (`<- chan`) yang siap. Lain halnya dengan
// _Buffered Channel_ akan menerima saja sejumlah value tanpa
// harus adanya penerima.
package main

import "fmt"

func main() {

	// Kita akan membuat `make` channel dengan type string
	// dan mem-buffer 2 value dalam channel tersebut.
	messages := make(chan string, 2)

	// Karena chanel ini di-_buffer_, maka kita bisa mengirimkan
	// value-value ini ke dalam channel tanpa penerima secara langsung.
	messages <- "buffered"
	messages <- "channel"

	// Kemudian kita bisa menerima dua value tersebut seperti biasa.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
