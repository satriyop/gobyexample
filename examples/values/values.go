// Go mempunyai bermacam-macam value types, antara lain
// strings, integers, floats, booleans, dan lain-lain.
// Berikut di bawah adalah beberapa contoh dasar.

package main

import "fmt"

func main() {

	// String, yang dapat digabungkan dengan `+`.
	fmt.Println("go" + "lang")

	// Integer dan float.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Boolean, dengan operator boolean.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
