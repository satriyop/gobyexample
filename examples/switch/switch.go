// _Statement switch_  mengekspresikan persyaratan
// dalam banyak percabangan.
package main

import (
	"fmt"
	"time"
)

func main() {

	// Berikut adalah `switch` sederhana
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Kita bisa menggunakan koma untuk memisahkan
	// beberapa _expression_ yang sama dalam satu
	// statement `case` yang sama. Kita juga menggunakan `default`
	// yang opsional pada contoh ini.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// `switch` tanpa _expression_ sebagai alternatif
	// logika if/else. Di sini kita menunjukkan _expression_ `case`
	// bisa saja bukan konstan.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// `switch` type akan membandingkan type bukan value.
	// Kita bisa menggunakan ini untuk mengetahui type dari
	// nilai interface. Dalam contoh ini, kita variabel `t` akan
	// mempunyai type yang sesuai dengan klausanya.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
