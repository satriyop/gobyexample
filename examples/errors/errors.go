// Mengkomunikasikan error secara eksplisit, dalam
// sebuah value yang di-return secara terpisah adalah
// hal yang disarankan (idiomatic) di Go. Hal ini
// berbeda dengan bahasa pemrograman lain seperti Java
// dan Ruby yang menggunakan _exception_ atau dengan bahasa
// C yang kadang meng-_overload_ hasil return (result/error).
// Pendekatan di Go, membuat error lebih mudah terlihat
// dimana sebuah fungsi akan mengembalikan error dan
// untuk meng-handle error tersebut menggunakan _construct_
// yang sama seperti yang digunakan untuk type selain error.

package main

import (
	"errors"
	"fmt"
)

// Berdasarkan kesepakatan, error adalah value return
// di posisi terakhir yang mempunyai type `error`, sebuah interface
// yang disediakan built-in.
func f1(arg int) (int, error) {
	if arg == 42 {

		// `errors.New` akan membuat sebuah value `error`
		// sederhana dengan pesan error yang bisa berikan
		// sebagai argument.
		return -1, errors.New("can't work with 42")

	}

	// Sebuah value `nil` pada posisi error mengindikasikan
	// bahwa tidak ada error.
	return arg + 3, nil
}

// Dimungkinkan untuk menggunakan type selain `error`
// dengan mengimplementasikan method `Error()` pada type
// tersebut. Di sini adalah variasi dari contoh di atas
// dengan menggunakan type custom yang secara eksplisit
// merepresentasikan sebuah error argumen.
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// Dalam contoh ini kita menggunakan sintaks
		// `&argError` untuk membentuk struct baru, memberikan
		// nilai pada field `arg` dan `prob`.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// Dua loop di bawah ini mengetes masing-masing fungsi
	// yang mengebalikan error. Perhatikan bahwa penggunakan
	// pengecekan error secara inline pada `if` adalah
	// hal yang disarankan di kode Go.
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// Jika kita ingin secara program ingin menggunakan data
	// pada custom error, kita akan membutuhkan error sebagai
	// sebuah instance dari type custom error dengan
	// menggunakan teknik _type assertion_.
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
