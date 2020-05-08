# Dengan menjalankan program ini akan terlihat bahwa
# mengatur state dengan goroutine menyelesaikan
# sekitar 80.000 total operasi.
$ go run stateful-goroutines.go
readOps: 71708
writeOps: 7177


# Dalam kasus ini, pendekatan dengan goroutine membutuhkan
# lebih banyak detail daripada menggunakan mutex. 
# Meskipun begituH hal ini berguna pada beberapa kasus
# sebagai contoh dimana kita mumpunyai channel lain yang 
# terlibat atau ketika mengatur beberapa mutex menjadikan
# program kita lebih rentan terhadap error. 
# Kita harus menggunakan  pendekatan mana yang 
# terasa lebih alami, terutama dalam  konteks 
# program yang kita buat.