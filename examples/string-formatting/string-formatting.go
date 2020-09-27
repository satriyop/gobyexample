// Go offers excellent support for string formatting in
// the `printf` tradition. Here are some examples of
// common string formatting tasks.

// Go menawarkan cara melakukan format string sesuai
// tradisi `printf`. Di sini adalah beberapa contoh
// dari melakukan format string.
package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	// Go offers several printing "verbs" designed to
	// format general Go values. For example, this prints
	// an instance of our `point` struct.

	// Go menawarkan beberapa `verbs` printing yang didesain
	// untuk melakukan format pada value Go yang umum. Sebagai
	// contoh, hal ini akan menampilkan instance dari struct
	// `point` yang sudah dideklarasikan.
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// If the value is a struct, the `%+v` variant will
	// include the struct's field names.

	// Jika value yang akan di-print adalah sebuah struct,
	// maka variasi `%+v` akan menampilkan nama nama field
	// pada struct tersebut.
	fmt.Printf("%+v\n", p)

	// The `%#v` variant prints a Go syntax representation
	// of the value, i.e. the source code snippet that
	// would produce that value.

	// Variasi `%#v` akan menampilkan representasi sintaks Go
	// dari value yang di-print sebagai contoh
	// kode sumber untuk menghasilkan value tersebut.
	fmt.Printf("%#v\n", p)

	// To print the type of a value, use `%T`.
	// Untuk menampilkan type dari value, gunakan `%T`.
	fmt.Printf("%T\n", p)

	// Formatting booleans is straight-forward.
	// Melakukan format pada boolean secara sederhana.
	fmt.Printf("%t\n", true)

	// There are many options for formatting integers.
	// Use `%d` for standard, base-10 formatting.

	// Ada beberapa opsi melakukan format pada integer.
	// Gunakan `%d` untuk menampilkan format standar base-10.
	fmt.Printf("%d\n", 123)

	// This prints a binary representation.
	// Ini akan menampilkan representasi biner.
	fmt.Printf("%b\n", 14)

	// This prints the character corresponding to the
	// given integer.

	// Ini akan menampilkan karakter yang sesuai dengan
	// nilai integer.
	fmt.Printf("%c\n", 33)

	// `%x` akan menampilkan encoding dalam hex.
	fmt.Printf("%x\n", 456)

	// There are also several formatting options for
	// floats. For basic decimal formatting use `%f`.

	// Ada juga beberapa opsi format untuk float.
	// untuk desimal sederhana, gunakan `%f`.
	fmt.Printf("%f\n", 78.9)

	// `%e` and `%E` format the float in (slightly
	// different versions of) scientific notation.

	// `%e` dan `%E` kan mem-format float dalam
	// notasi scientifik yang sedikit berbeda.
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// For basic string printing use `%s`.
	// Untuk print string sederhana, gunakan `%s`.
	fmt.Printf("%s\n", "\"string\"")

	// To double-quote strings as in Go source, use `%q`.
	// Untuk menampilkan format double-quote string
	// gunakan `%q`.
	fmt.Printf("%q\n", "\"string\"")

	// As with integers seen earlier, `%x` renders
	// the string in base-16, with two output characters
	// per byte of input.

	// Sebagai mana integer yang kita lihat di awal, `%x`
	// akan menampilkan string dalam base-16, dengan dua output
	// karakter per byte dari input.
	fmt.Printf("%x\n", "hex this")

	// To print a representation of a pointer, use `%p`.

	// Untuk menampilkan representasi dari sebuah pointer,
	// gunakan `%p`.
	fmt.Printf("%p\n", &p)

	// When formatting numbers you will often want to
	// control the width and precision of the resulting
	// figure. To specify the width of an integer, use a
	// number after the `%` in the verb. By default the
	// result will be right-justified and padded with
	// spaces.

	// Ketika melakukan format angka, kita kadang ingin
	// mengontrol lebar dan presisi dari hasilnya.
	// untuk megatur lebar dari integer, gunakan sebuah
	// angka setelah tanda `%`. Secara default hasilnya
	// akan rata kanan dan ditambahkan padding.
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// You can also specify the width of printed floats,
	// though usually you'll also want to restrict the
	// decimal precision at the same time with the
	// width.precision syntax.

	// Kita bisa mengatur lebar dari angka float
	// yang diprint, meskipun kita juga ingin mengatur
	// presisi desimal pada saat yang bersamaan dengan
	// sintaks width.precision.
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// To left-justify, use the `-` flag.

	// Untuk rata kiri, gunakan flag `-`.
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// You may also want to control width when formatting
	// strings, especially to ensure that they align in
	// table-like output. For basic right-justified width.

	// Kita juga bisa mengatur lebar ketika melakukan
	// format pada string, terutama agak terlihat
	// seperti table. Untuk pengaturan rata kanan sederhana.
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// To left-justify use the `-` flag as with numbers.

	// Untuk rata kiri, gunakan flag `-` sebagaimana dengan angka.
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// So far we've seen `Printf`, which prints the
	// formatted string to `os.Stdout`. `Sprintf` formats
	// and returns a string without printing it anywhere.

	// kita telah melihat `Printf`, yang menampilkan
	// hasil ke `os.Stdout`. `Sprintf` akan melakukan
	// format dan mengembalikan nilai string tanpa menampilkan
	// kemana saja.
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// You can format+print to `io.Writers` other than
	// `os.Stdout` using `Fprintf`.

	// Kita bisa melakukan format+print pada `io.Writer`
	// selain ke `os.Stdout` menggunakan `Fprintf`.
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
