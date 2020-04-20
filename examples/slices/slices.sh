# Perhatikan bahwa type slice berbeda dengan array,
# Slice ditampilkan mirip dengan array oleh 
# fungsi `fmt.Prinln`.
$ go run slices.go
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]

# Silakan cek ini [postingan blog bagus](http://blog.golang.org/2011/01/go-slices-usage-and-internals.html)
# ditulis oleh tim Go untuk mengetahui yang detail 
# tentang desain & implementasi slice di Go.

# Sekarang kita sudah melihat array dan slice, 
# selanjutnya mari kita lihat 
# struktur data penting lainnya di Go: maps.
