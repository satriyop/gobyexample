# Dengan menjalankan akan terlihat bahwa kita mengeksekusi
# sekitar 90.000 operasi terhadap `state` yang kita 
# sinkronisasi dengan `mutex`.
$ go run mutexes.go
readOps: 83285
writeOps: 8320
state: map[1:97 4:53 0:33 2:15 3:2]

# Selanjutnya kita akan melihat bagaimana 
# mengimplementasikan pengaturan state 
# yang sama dengan ini dengan hanya menggunakan
# goroutine dan channel.