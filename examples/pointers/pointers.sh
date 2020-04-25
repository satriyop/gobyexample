# `zeroval` tidak mengubah nilai `i` dalam `main`, 
# sedangkan `zeroptr` mengubahnya karena menggunakan 
# referensi alamat memori pada variabel 'i' tersebut.
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
