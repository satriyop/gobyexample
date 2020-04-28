// Go mendukung _method_ yang didefinisikan pada type struct.
package main

import "fmt"

type rect struct {
	width, height int
}

// Method `area` mempunyai _receiver_ yaitu type `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// Method bisa didefinisikan baik untuk  type dengan jenis _pointer_ maupun _value_.
// Berikut adalah contoh dengan _receiver_ dalam bentun value.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Di sini kita memanggail 2 method yang sudah didefinisikan pada struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go secara otomatis menangani konversi antara
	// value dan pointer dalam pemanggilan sebuah method.
	// Kita bisa menggunakan receiver dengan type pointer
	// untuk menghindari proses copy saat method dipanggil
	// atau dalam hal kita ingin mengubah data di structnya.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
