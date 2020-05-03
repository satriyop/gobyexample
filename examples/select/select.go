// _Select_ di Go, memungkinkan kita untuk menunggu
// pada beberapa operasi / eksekusi kode dengan channel.
// Kombinasi dari goroutine, channel, dan select
// merupakan fitur yang hebat.
package main

import (
	"fmt"
	"time"
)

func main() {

	// Dalam contoh ini, kita akan melakukan _select
	// pada dua channel.
	c1 := make(chan string)
	c2 := make(chan string)

	// Setiap channel akan menerima value setelah sejumlah
	// waktu sebagai simulasi, misalnya  dari eksekusi dari goroutine
	// yang _concurrent_ dari operasi blocking RPC.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Kita akan menggunakan `select` untuk menunggu
	// kedua value tersebut secara bersamaan, dan
	// menampilkannya ketika diterima.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
