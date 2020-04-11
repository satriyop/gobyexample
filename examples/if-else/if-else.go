// Percabangan dengan `if` dan `else` di Go
// adalah hal yang mudah.
package main

import "fmt"

func main() {

	// Berikut adalah contoh sederhana.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Kita bisa menggunakan statement `if` tanpa else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Sebuah statement dapat disisipkan didepan percabangan;
	// apapun variabel yang dideklarasikan akan tersedia di semua
	// percabangan.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Note that you don't need parentheses around conditions
// in Go, but that the braces are required.
