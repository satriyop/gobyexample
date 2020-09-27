// Gunakan `os.Exit` untuk segera mengakhiri
// dan menambahkan status exit pada program.
package main

import (
	"fmt"
	"os"
)

func main() {

	// `defer` tidak akan dijalankan ketika kita menggunakan
	// `os.Exit`, jadi perintah `fmt.Println` ini tidak akan
	// pernah dipanggil.
	defer fmt.Println("!")

	// Mengakhiri program dengan status 3.
	os.Exit(3)
}

// Perhatikan bahwa tidak seperti misalnya C, Go tidak
// menggunakan nilai return integer dari `main` untuk
// mengindikasikan sebuat status exit. Jika kita ingin
// mengakhiri program dengan status selain 0, maka kita
// harus menggunakan `os.Exit`.
