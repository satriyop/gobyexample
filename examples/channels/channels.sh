# Ketika kita menjalankan program, value "ping" 
# diterima oleh goroutine lain melalui channel yang 
# kita buat.
$ go run channels.go 
ping

# Secara default proses mengirim dan menerima 
# akan melakukan _block_ sampai pengirim 
# dan penerima siap. 
# Hal ini membuat kita menunggu pada akhir program untuk 
# menerima "ping" tanpa harus menggunakan 
# proses sinkronisasi lainnya.