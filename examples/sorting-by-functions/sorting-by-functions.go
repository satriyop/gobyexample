// Kadang kita ingin melakukan pengurutan pada sebuah
// collection bukan dari urutan alaminya. Sebagai contoh
// Kita ingin mengurutkan string dari panjang bukan
// secara alfabet. Dalam contoh ini kita akan melakukan
// pengurutan terkustomisasi di Go.
package main

import (
	"fmt"
	"sort"
)

// Untuk melakukan pengurutan oleh sebuah fungsi kustomisasi,
// kita membutuhkan type yang berkesesuaian. Di sini kita
// membuat type `byLength` yang akan menjadi alias dari
// type builtin `[]string`.
type byLength []string

// Kita mengimplementasikan `sort.Interface` - `Len`, `Less`,
// dan `Swap` - pada type kita sehingga kita bisa menggunakan
// fungsi generic `Sort` pada package `sort`.
// Method `Len` dan `Swap` biasanya akan mirip pada semua type sedangkan
// `Less` akan menyimpan logika pengurutan yang telah dikustomisasi.
// Dalam contoh kita, kita ingin mengurutkan berdasarkan panjang
// sehingga kita menggunakan `len(s[i])` dan `len(s[j])` di sini.
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// Dengan semua persiapan di atas, saatnya kita
// menggunakan pengurutan terkustomisasi kita dengan
// mengubah slice `fruits` menjadi `byLength` dan kemudian
// menggunakan `sort.Sort` pada type tersebut.
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
