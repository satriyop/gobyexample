// Go mendukung _constants_ dari character, string,
// boolean, dan angka numerik.
package main

import (
	"fmt"
	"math"
)

// `const` mendeklarasikan sebuah constant.
const s string = "constant"

func main() {
	fmt.Println(s)

	// statement `const` dapat ditulis dimana saja
	// sebagai mana statement `var`.
	const n = 500000000

	// Constant expressions perform arithmetic with
	// arbitrary precision.
	// Constant expression akan melakukan perhitungan
	// dengan presisi yang arbitrary
	const d = 3e20 / n
	fmt.Println(d)

	// Constant numerik tidak mempunyai type
	// sampai diberikan,dengan konversi ekplisit
	// contohnya.
	fmt.Println(int64(d))

	// Sebuah angka dapat diberikan sebuah type
	// dengan menggunakannya dalam sebuah context yang
	// membutuhkannya, seperti variable assignment
	// atau function call. Dalam contoh ini,
	// kita `math.Sin` akan mengharapkan `float64`.
	fmt.Println(math.Sin(n))
}
