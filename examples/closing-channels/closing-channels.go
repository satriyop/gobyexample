// Menutup sebuah channel mengindikasikan bahwa tidak ada
// lagi value yang akan dikirimkan pada channel tersebut.
// Hal ini berguna untuk mengkomunikasikan bahwa proses
// telah selesai di sisi penerima channel.
package main

import "fmt"

// Dalam contoh ini, kita menggunakan sebuah channel
// `jobs` yang akan mengkomunikasikan pekerjaan yang harus
// diselesaikan dari gorouting `main()`.
// Ketika tidak ada lagi jobs untuk _worker_ kita akan menutup
// channel `jobs`.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Ini adalah goroutine dari _worker_. goroutine ini akan
	// secara berulang menerima dari value dari channel `jobs`
	// dalam sintaks `j, more := <- jobs`.
	// Dalam bentuk penerima 2-value  yang khusus ini,
	// value dari `more` akan `false` ketika
	// `jobs` telah ditutup dan semua value dalam channel tersebut
	// sudah diterima.
	// Kita menggunakan ini untuk memberitahu pada channel `done`
	// ketika kita sudah melakukan semua proses(jobs).
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Kode ini akan mengirimkan 3 value ke channel jobs
	// kepada worker, dan menutup channel tersebut.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// Kita menunggu worker menggunakan
	// pendekatan [synchronization](channel-synchronization)
	// seperti yang sudah kita lakukan sebelumnya.
	<-done
}
