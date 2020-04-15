// _Slice_ adalah tipe data yang utama di Go,
// memberikan _interface_ ke urutan yang lebih
// hebat daripada array.
package main

import "fmt"

func main() {

	// Tidak seperti array, slice hanya di-type
	// oleh elemen yang ada didalamnya (bukan oleh jumlah dari elemen).
	// Untuk membuat slice kosong dengan _non-zero length_,
	// gunakan fungsi builtin `make`. Di contoh ini kita membuat slice
	// dari `string` dengan panjang `3` (diinisiasikan _zero-valued_
	// untuk nilai elemennya).
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// Kita bisa mengatur dan mengakses elemen seperti dalam array.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `fungsi len` akan mengembalikan panjang dari slice.
	fmt.Println("len:", len(s))

	// Sebagai tambahan dari operasi dasar ini, slice
	// mendukung beberapa operasi lain yang membuatnya lebih
	// kaya daripada array. Salah satunya adalah `append`, yang
	// akan mengambalikan sebuah slice yang mempunyai satu / lebih
	// nilai. Perhatikan bahwa kita harus menerima nilai yang dikembalikan
	// dari `append` karena kita akan mendapatkan sebuah slice yang baru.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slice juga bisa di `copy`. Di sini kita membuat sebuah
	// slice kosong `c` yang mempunyai ukuran yang sama dengan `s`
	// dan kita copy `s` ke dalam `c`.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slice mendukung sebuah operator "slice" dengan sintaks
	// `slice[low:high]`. Dalam contoh ini, kita akan mendapatkan
	// elemen `s[2]`, `s[3]`, and `s[4]`.
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Contoh ini akan men-slice sampai `s[5]` (elemen 5 tidak termasuk).
	l = s[:5]
	fmt.Println("sl2:", l)

	// Contoh ini akan men-slice sampai `s[2]` (elemen 2 termasuk).
	l = s[2:]
	fmt.Println("sl3:", l)

	// Kita bisa mendeklarasikan dan menginisiasikan
	// sebuah variabel untuk slice dalam satu baris.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slice bisa disusun menjadi struktur data multi-dimensi.
	// Ukuran dari _inner slice_ bisa bervariasi, tidak seperti
	// array multi-dimensi.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
