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
	// Go menawarkan beberapa `verbs` printing yang didesain
	// untuk melakukan format pada value Go yang umum. Sebagai
	// contoh, hal ini akan menampilkan instance dari struct
	// `point` yang sudah dideklarasikan.
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// Jika value yang akan di-print adalah sebuah struct,
	// maka variasi `%+v` akan menampilkan nama nama field
	// pada struct tersebut.
	fmt.Printf("%+v\n", p)

	// Variasi `%#v` akan menampilkan representasi sintaks Go
	// dari value yang di-print sebagai contoh
	// kode sumber untuk menghasilkan value tersebut.
	fmt.Printf("%#v\n", p)

	// Untuk menampilkan type dari value, gunakan `%T`.
	fmt.Printf("%T\n", p)

	// Melakukan format pada boolean secara mudah.
	fmt.Printf("%t\n", true)

	// Ada beberapa opsi melakukan format pada integer.
	// Gunakan `%d` untuk menampilkan format standar base-10.
	fmt.Printf("%d\n", 123)

	// Ini akan menampilkan representasi biner.
	fmt.Printf("%b\n", 14)

	// Ini akan menampilkan karakter yang sesuai dengan
	// nilai integer.
	fmt.Printf("%c\n", 33)

	// `%x` akan menampilkan encoding dalam hex.
	fmt.Printf("%x\n", 456)

	// Ada juga beberapa opsi format untuk float.
	// untuk desimal sederhana, gunakan `%f`.
	fmt.Printf("%f\n", 78.9)

	// `%e` dan `%E` kan mem-format float dalam
	// notasi scientifik yang sedikit berbeda.
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// Untuk print string sederhana, gunakan `%s`.
	fmt.Printf("%s\n", "\"string\"")

	// Untuk menampilkan format double-quote string
	// gunakan `%q`.
	fmt.Printf("%q\n", "\"string\"")

	// Sebagai mana integer yang kita lihat di awal, `%x`
	// akan menampilkan string dalam base-16, dengan dua output
	// karakter per byte dari input.
	fmt.Printf("%x\n", "hex this")

	// Untuk menampilkan representasi dari sebuah pointer,
	// gunakan `%p`.
	fmt.Printf("%p\n", &p)

	// Ketika melakukan format angka, kita kadang ingin
	// mengontrol lebar dan presisi dari hasilnya.
	// untuk megatur lebar dari integer, gunakan sebuah
	// angka setelah tanda `%`. Secara default hasilnya
	// akan rata kanan dan ditambahkan padding.
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// Kita bisa mengatur lebar dari angka float
	// yang diprint, meskipun kita juga ingin mengatur
	// presisi desimal pada saat yang bersamaan dengan
	// sintaks width.precision.
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// Untuk rata kiri, gunakan flag `-`.
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// Kita juga bisa mengatur lebar ketika melakukan
	// format pada string, terutama agak terlihat
	// seperti table. Untuk pengaturan rata kanan sederhana.
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// Untuk rata kiri, gunakan flag `-` sebagaimana dengan angka.
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// kita telah melihat `Printf`, yang menampilkan
	// hasil ke `os.Stdout`. `Sprintf` akan melakukan
	// format dan mengembalikan nilai string tanpa menampilkan
	// kemana saja.
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// Kita bisa melakukan format+print pada `io.Writer`
	// selain ke `os.Stdout` menggunakan `Fprintf`.
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
