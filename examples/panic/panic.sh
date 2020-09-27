# Jika kita menjalankan program ini, maka akan terjadi
# panic, dan menampilkan pesan error dan 
# trace dari goroutine,
# kemudian program akan berakhir dengan non-zero status.
$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# Perhatikan bahwa tidak seperti beberapa bahasa 
# pemrograman lain yang menggunakan exception untuk 
# menangani berbagai error,
# di Go adalah hal yang disarankan untuk menggunakan 
# value return yang mengindikasikan sebuah error 
# bila memungkinkan.