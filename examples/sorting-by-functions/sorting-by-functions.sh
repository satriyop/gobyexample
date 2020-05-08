# Dengan menjalankan program terlihat sebuah daftar
# yang diurutkan berdasarkan panjangnya sebagaimana
# yang diharapkan.
$ go run sorting-by-functions.go 
[kiwi peach banana]


# Dengan mengikuti pola ini : membuat type terkustomisasi,
# mengimplementasikan tiga  method `interface` pada type 
# tersebut, dan memanggil fungsi sort.Sort pada collection
# type kita, kita bisa mengurutkan slice dengan 
# fungsi yang sesuai kemauan kita.