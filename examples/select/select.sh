# Kita menerima value `"one"` dan kemudian `"two"`
# seperti yang kita harapkan.
$ time go run select.go 
received one
received two

# Perhatikan bahwa total waktu eksekusi hanya sekitar
# ~2 detik karena kedua perintah `Sleep` dieksekusi
# secara concurrent.
real	0m2.245s
