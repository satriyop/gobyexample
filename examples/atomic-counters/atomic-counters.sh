# Kita mengharapkan mendapatkan persis 50.000 operasi.
# Jika kita menggunakan non-atomic `ops++` untuk 
# menaikkan counter kemungkinan besar kita akan 
# mendapatkan jumlah lain yang berbeda setiap kali
# menjalankan program karena goroutine akan mempengaruhi
# satu sama lain. Kita akan mendapatkan data race failure
# bila kita menjalankan program dengan flag `-race`.
$ go run atomic-counters.go
ops: 50000

# Selanjutnya kita akan melihat mutex, 
# tool lainnya yang untuk mengatur state.