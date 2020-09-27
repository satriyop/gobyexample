// Go offers built-in support for [regular expressions](http://en.wikipedia.org/wiki/Regular_expression).
// Here are some examples of  common regexp-related tasks
// in Go.

// Go mendukung fitur [regular expressions](http://en.wikipedia.org/wiki/Regular_expression) secara built-in.
// Berikut adalah beberapa contoh penggunaan regexp di Go.
package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// This tests whether a pattern matches a string.

	// Untuk menguji bahwa sebuah pola cocok pada sebuah string.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Above we used a string pattern directly, but for
	// other regexp tasks you'll need to `Compile` an
	// optimized `Regexp` struct.

	// Pada contoh di atas kita menggunakan pola pada
	// string secara langsung, tapi untuk penggunaan
	// regexp yang lain kita perlu untuk melakukan `Compile`
	// struct `Regexp` yang telah dioptimasi.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Many methods are available on these structs. Here's
	// a match test like we saw earlier.

	// Berbagai method tersedia pada struct ini.
	// Di sini kita menguji pola seperti contoh di atas.
	fmt.Println(r.MatchString("peach"))

	// This finds the match for the regexp.

	// Hal ini akan menemukan kecocokan pada regexp.
	fmt.Println(r.FindString("peach punch"))

	// This also finds the first match but returns the
	// start and end indexes for the match instead of the
	// matching text.

	// Hal ini juga akan menemukan kecocokan pertama tapi
	// mengembalikan index awal dan akhir untuk kecocokan
	// bukan text yang cocok.
	fmt.Println(r.FindStringIndex("peach punch"))

	// The `Submatch` variants include information about
	// both the whole-pattern matches and the submatches
	// within those matches. For example this will return
	// information for both `p([a-z]+)ch` and `([a-z]+)`.

	// Variasi `Submatch` menyertakan informasi
	// baik tentang kecocokan pola secara keseluruhan dan
	// kecocokan sebagian (submatch). Contoh di bawah ini akan
	// mengembalikan informasi baik `p([a-z]+)ch`
	// maupun  `([a-z]+)`.
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Similarly this will return information about the
	// indexes of matches and submatches.
	// Hal yang sama, contoh di bawah ini akan mengembalikan
	// informasi tentang kecocokan index dan submatch-nya.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// The `All` variants of these functions apply to all
	// matches in the input, not just the first. For
	// example to find all matches for a regexp.
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// These `All` variants are available for the other
	// functions we saw above as well.
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// Providing a non-negative integer as the second
	// argument to these functions will limit the number
	// of matches.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// Our examples above had string arguments and used
	// names like `MatchString`. We can also provide
	// `[]byte` arguments and drop `String` from the
	// function name.
	fmt.Println(r.Match([]byte("peach")))

	// When creating constants with regular expressions
	// you can use the `MustCompile` variation of
	// `Compile`. A plain `Compile` won't work for
	// constants because it has 2 return values.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// The `regexp` package can also be used to replace
	// subsets of strings with other values.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// The `Func` variant allows you to transform matched
	// text with a given function.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
