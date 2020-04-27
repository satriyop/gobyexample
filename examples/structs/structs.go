// Di bahasa pemrograman Go, _struct_ adalah sebuah _type_
// yang berupa kumpulan dari _field-field_.
// Struct berguna untuk mengelompokkan data-data
// sehingga bisa disimpan.
package main

import "fmt"

// Type dari _struct_ `person` ini mempunyai _field_ `name` dan `age`.
type person struct {
	name string
	age  int
}

// Fungsi NewPerson ini akan menbentuk _struct_ person dengan argumen name.
func NewPerson(name string) *person {
	// Kita bisa mengembalikan sebuah pointer pada variabel lokal
	// karena variabel lokal ini akan masih dalam _scope_ fungsi.
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// Sintaks ini akan menciptakan sebuah struct baru.
	fmt.Println(person{"Bob", 20})

	// Kita bisa menuliskan nama field ketika menginisiasi
	// sebuah struct.
	fmt.Println(person{name: "Alice", age: 30})

	// Field yang tidak diberi nilai akan `zero-valued`
	// (yang nilainya akan tergantung dengan tipenya)
	fmt.Println(person{name: "Fred"})

	// Prefiks `&` akan menghasilkan sebuah pointer
	// terhadap struct tersebut.
	fmt.Println(&person{name: "Ann", age: 40})

	// Mengenkapsulasi pembuatan struct baru dalam sebuah fungsi _constructor_
	// disarankan (idiomatic).
	fmt.Println(NewPerson("Jon"))

	// Gunakan dot (.) untuk mengakses field dari sebuah struct.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// Kita juga bisa menggunakan dot dalam pointer struct.
	// Pointer tersebut akan secara otomatis di-_dereference_.
	sp := &s
	fmt.Println(sp.age)

	// Struct merupakan tipe data yang _mutable_ (bisa diubah).
	sp.age = 51
	fmt.Println(sp.age)
}
