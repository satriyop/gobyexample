// Sebuah `panic` bisa diartikan telah terjadi sesuatu
// yang tidak diharapkan. Seringnya kita menggunakan panic
// saat terjadi error yang seharusnya tidak terjadi
// bila semua operasi berjalan normal, atau kita tidak
// menyiapkan penanganan error tersebut secara baik.
package main

import "os"

func main() {

	// Kita akan menggunakan panic sepanjang situs ini
	// untuk mengecek error yang tidak terduga. Contoh
	// ini adalah satu-satunya program di situs ini yang
	// didesain untuk panic.
	panic("a problem")

	// Penggunaan panic pada umumnya adalah untuk membatalkan
	// jika sebuah fungsi mengembalikan sebuah value error yang mana
	// kita tidak tahu  bagaimana cara menanganinya (atau tidak mau).
	// Di contoh ini adalah `panic` bila terjadi error saat membuat
	// sebuah file baru.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
