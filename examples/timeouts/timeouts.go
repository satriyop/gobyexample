// _Timeout_ adalah hal penting bagi program yang
// menghubungkan _resource_ eksternal atau hal lain yang
// membutuhkan waktu lama pada eksekusi. Mengimplementasikan
// timeout pada Go menjadi mudah dan elegan dengan menggunakan
// channel dan `select`.
package main

import (
	"fmt"
	"time"
)

func main() {

	// Misalnya kita akan mengeksekusi sebuah _external call_
	// yang akan mengembalikan hasilnya pada channel `c1`
	// setelah 2 detik. Perhatikan ahwa channel ini di-_buffer_,
	// sehingga pengiriman pada goroutine ini sifatnya non-blocking.
	// hal ini adalah pola yang umum dipakai untuk menghindari
	// gorouting _leaks_ pada kasus dimana channel tidak pernah dibaca.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Di contoh ini, kita menggunakan `select` yang
	// mengimplementasikan timeout.
	// Sintask `res := <-c1` akan menunggu dari hasil
	// dan sintaks `<-Time.After` akan menunggu value
	// yang akan dikirimkan setelah waktu timeout 1 detik.
	// Karena `select` akan memilih yang terlebih dahulu
	// menerima value, maka kita akan mengeksekusi
	// case timeout, hanya bila operasi ini berjalan lebih
	// dari 1 detik seperti yang kita atur di kode.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// Jika kita mengijinkan waktu timeout lebih lama dari 3 detik,
	// maka penerima dari `c2` akan sukses dan menampilkan hasilnya.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
