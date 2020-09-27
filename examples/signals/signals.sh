# Ketika kita menjalankan program ini, program akan
# menunggu sebuah sinyal. Dengan menekan `ctrl-C` (pada
# terminal ditampilkan dengan `^C`) kita bisa mengirimkan
# sinyal `SIGINT`, yang akan mengakibatkan program untuk 
# menampilkan `interrupt` dan akhirnya mengakhiri program.
$ go run signals.go
awaiting signal
^C
interrupt
exiting
