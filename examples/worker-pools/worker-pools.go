// Dalam contoh ini kita akan mengimplementasikan
// sebuah _worker pool_ menggunakan goroutine dan channel.
package main

import (
	"fmt"
	"time"
)

// Kita akan menjalankan beberapa instance concurrent pada
// worker seperti di contoh ini. Worker-worker ini akan menerima
// tugas channel `jobs` dan mengirimkan hasilnya pada channel `result`.
// Kita akan melakukan _sleep_ selama satu detik setiap tugas untuk
// men-simulasikan tugas berat atau yang membutuhkan waktu.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// Untuk menggunakan kumpulan dari worker tersebut kita
	// harus mengirimkan tugas dan mengumpulkan hasil tugasnya.
	// Kita menggunakan 2 channel untuk itu.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Kode ini akan memulai 3 worker, awalnya akan
	// membuat block karena belum ada tugas yang dikirimkan.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Di sini kita mengirimkan 5 `jobs` kemudian menutupnya
	// dengan `close` untuk mengindikasikan bahwa pekerjaan
	// telah selesai.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Akhirnya kita mengumpulkan semua hasil pekerjaan.
	// Hal ini juga memastikan bahwa goroutine worker
	// telah selesai. Sebagai metode lain untuk menunggu
	// goroutine selesai adalah menggunakan [WaitGroup](waitgroups).
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
