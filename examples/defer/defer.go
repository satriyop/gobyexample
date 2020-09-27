// _Defer_ digunakan untuk memastikan bahwa sebuah
// pemanggilan sebuah fungsi akan dilakukan dikemudian
// dalam sebuah eksekusi program, biasanya digunakan
// dalam rangka bersih-bersih. `defer` sering digunakan
// dimana kata kunci `ensure` dan `finally` digunakan
// dalam bahasa pemrograman lain.
package main

import (
	"fmt"
	"os"
)

// Misalnya kita ingin membuat sebuah file, dan
// menuliskan ke dalam file tersebut, dan kemudian
// menutup file tersebut ketika sudah selesai.
// Maka dengan `defer` kita bisa melakukannya seperti ini.
func main() {

	// Segera setelah object file tercipta dengan
	// `createFile`, kita men-`defer` penutupan file
	// dengan memanggil `closeFile`
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// Adalah penting untuk mengecek bila terjadi error
	// ketika menutup file, meskipun dalam sebuah fungsi
	// yang di-defer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
