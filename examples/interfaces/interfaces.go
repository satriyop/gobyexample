// _Interface_ adalah kumpulan dari _signature method_
// (definisi dari method/fungsi, tanpa body) yang mempunyai nama.
package main

import (
	"fmt"
	"math"
)

// Berikut ini adalah interface sederhana untuk
// yang bernama geometry
type geometry interface {
	area() float64
	perim() float64
}

// Dalam contoh ini kita akan mengimplementasikan
// interface di ada type `rect` dan `circle`.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Untuk mengimplementasikan interface di Go, kita hanya
// perlu mengimplementasikan semua method yang ada
// di interface tersebut. Dalam contoh ini, kita mengimplementasikan
// interface `geometry` pada type `rect`.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Implementasi interface geometry pada type `circle`.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Jika dalam sebuah variabel mempunyai type interface,
// maka kita bisa memanggil method yang ada di interface tersebut.

// Di sini sebuah fungsi biasa `measure` menggunakan type interface `geometry`
// sebagai parameter artinya apapun type yang mengimplementasikan interface `geometry`
// seperti `rect` atau `circle` bisa dipakai digunakan sebagai argumen saat pemanggilan fungsi.
// Dengan begitu fungsi ini bisa memanggil method-method yang ada di interface `geometry` terhadap
// type yang diberikan.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// type struct `circle` dan `rect` mengimplementasikan
	// interface `geometry` sehingga kita bisa menggunakan
	// _instance_ dari struct tersebut sebagai argumen pada
	// fungsi `measure`.
	measure(r)
	measure(c)
}
