# Ketika kita menjalankan program ini, kita melihat
# output dari panggilan terhadap fungsi 
# yang menge-block terlebih dulu, lalu menunjukkan output
# yang selang seling antara dua goroutine, yang 
# menunjukkan bahwa goroutine dijalankan secara
# concurrent oleh runtime Go.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done


# Selanjutnya kita akan melihat channel, yang bisa
# menjadi pelengkap gorouting pada program Go
# yang menggunakan fitur concurrent.