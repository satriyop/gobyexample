# Jika kita menjalankan `exit.go` menggunakan
# `go run`, pada saat program berakhir maka
# status exit akan ditampilkan di layar.
$ go run exit.go
exit status 3


# Dengan melakukan build dan mengeksekusi binary
# maka kita bisa melihat status exit di terminal.
$ go build exit.go
$ ./exit
$ echo $?
3

# Perhatikan bahwa `!` dari program 
# tidak pernah ditampilkan.