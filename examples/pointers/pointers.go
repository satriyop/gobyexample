// Go mendukung <em><a href="http://en.wikipedia.org/wiki/Pointer_(computer_programming)">pointers</a></em>,
// sehingga kita bisa melakukan _pass reference_
// pada sebuah value dan menyimpannya dalam program kita.
package main

import "fmt"

// Kita akan menunjukkan perbedaan  _pointer_ dengan _value_
// dalam dua fungsi : `zeroval` dan `zeroptr`.
// `zeroval` mempunyai parameter `int`, sehingga argumen
// yang akan diberikan ketika fungsi ini dipanggil
// harus sebuah _value_. Fungsi `zeroval` dalam hal ini
// kan mendapatkan sebuah salinan dari `ival` yang berbeda
// dari vairabel yang digunakan sebagai argumen
// saat fungsi ini dipanggil.
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` mempunyai parameter `*int`, yang berarti
// bahwa fungsi ini akan menerima sebuah pointer dari `int`.
// `*iptr` yang ada di badan fungsi ini akan melakukan
// _dereference_ dari pointer tersebut dari alamat memori
// sehingga mendapatkan value dari alamat memori tersebut.
// Memberikan sebuah vaue pada sebuah _deferenced pointer_ akan
// mengubah value pada alamat memori tersebut.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// sintaks `&i` akan memberikan alamat memori dari `i`
	// sebagai contoh sebuah pointer dari `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointer juga dapat ditampilkan seperti ini.
	fmt.Println("pointer:", &i)
}
