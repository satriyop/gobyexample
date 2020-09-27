// The standard library's `strings` package provides many
// useful string-related functions. Here are some examples
// to give you a sense of the package.

// Pustaka standar `strings` menyeiakan berbagai fungsi
// yang berkaitan dengan string. Di sini beberapa contoh
// akan menunjukkan penggunakan dari package ini.
package main

import (
	"fmt"
	s "strings"
)

// We alias `fmt.Println` to a shorter name as we'll use
// it a lot below.

// Kita meng-aliaskan `fmt.Println` ke dalam nama yang
// lebih pendek karena kita akan sering menggunakannya.
var p = fmt.Println

func main() {

	// Here's a sample of the functions available in
	// `strings`. Since these are functions from the
	// package, not methods on the string object itself,
	// we need pass the string in question as the first
	// argument to the function. You can find more
	// functions in the [`strings`](http://golang.org/pkg/strings/)
	// package docs.

	// Di sini adalah beberapa contoh fungsi yang ada
	// pada `strings`. Karena fungsi-fungsi ini berasal
	// dari package, bukan merupakan method dari object
	// string sendiri, kita perlu memberikan string yang
	// akan dioperasikan sebagai argumen pertama pada fungsi.
	// Kita bisa membaca lebih lengkap pada dokumentasi
	// package [`strings`](http://golang.org/pkg/strings/).
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	// Not part of `strings`, but worth mentioning here, are
	// the mechanisms for getting the length of a string in
	// bytes and getting a byte by index.

	// Bukan merupakan bagian dari  package `strings`, tapi
	// cukup penting disebutkan disini adalah mekanisme untuk
	// mendapatkan panjang dari sebuah string dalam bytes dan
	// mendapatkan sebuah byte dari index.
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}

// Note that `len` and indexing above work at the byte level.
// Go uses UTF-8 encoded strings, so this is often useful
// as-is. If you're working with potentially multi-byte
// characters you'll want to use encoding-aware operations.
// See [strings, bytes, runes and characters in Go](https://blog.golang.org/strings)
// for more information.

// Perhatikan bahwa `len` dan indexing di atas bekerja pada
// level byte. Go menggunakan UTF-8 encoded string, sehingga
// bisa berjalan seperti yang diharapkan. Jika kita bekerja dengan
// karater yang mungkin multi-byte, maka kita perlu menggunakan operasi
// yang memperhatikan encoding.
// Baca [strings, bytes, runes and characters in Go](https://blog.golang.org/strings)
// untuk informasi lebih jauh.
