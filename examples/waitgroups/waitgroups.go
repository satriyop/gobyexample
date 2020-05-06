// Untuk menunggu beberapa goroutine selesai,
// kita bisa menggunakan _wait group_
package main

import (
	"fmt"
	"sync"
	"time"
)

// Ini adalah fungsi yang akan berjalan pada setiap
// goroutine. Perhatikan bahwa sebuah WaitGroup harus
// diberikan sebagai pointer pada parameter fungsi.
func worker(id int, wg *sync.WaitGroup) {
	// Akan men-decrement counter WaitGroup,
	// sehingga saat return, akan memberitahu WaitGroup bahwa
	// goroutine telah selesai.
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	// Menggunakan _sleep_ untuk simulasi pekerjaan
	// yang membutuhkan waktu untuk selesai.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// WaitGroup ini digunakan untuk menunggu
	// semua goroutine yang sudah berjalan untuk
	// selesai.
	var wg sync.WaitGroup

	// Menjalankan beberapa goroutine dan menambahkan counter
	// dari WaitGroup untuk setiap goroutinenya.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Melakukan block sampai counter dari WaitGroup
	// kembali ke 0; Semua worker akan tahu kalau mereka
	// telah selesai.
	wg.Wait()
}
