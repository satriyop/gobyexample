// Di Go, _variables_ dideklarasikan secara eksplisit
// dan digunakan oleh compiler mengecek type dari
// function calls

package main

import "fmt"

func main() {

	// `var` mendeklarasi sebuah variabel atau lebih.
	var a = "initial"
	fmt.Println(a)

	// Mendeklarasikan lebih dari satu variabel sekaligus.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go akan menyimpulkan type dari variabel yang diinisiasi.
	var d = true
	fmt.Println(d)

	// Variabel-variabel yang dideklarasikan tanpa
	// diinisiasikan adalah _zero-value_ . Sebagai contoh,
	// zero value dari sebuah `int` adalah 0.
	var e int
	fmt.Println(e)

	// sintaks := adalah cara singkat untuk mendeklarasikan
	// dan menginisiasi sebuah variabel, sebagai contoh
	// `var f string = "apple"` dalam hal ini dapat dituliskan
	// seperti ini.
	f := "apple"
	fmt.Println(f)
}
