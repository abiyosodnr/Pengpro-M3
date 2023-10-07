package main

import "fmt"

func main() {
	// var a, b, x int
	// var hasil bool
	// fmt.Scan(&a, &b)

	// x = a - b
	// hasil = (x%2 != 0)

	// fmt.Println(hasil)

	//soalNomor1

	// const pi float64 = 22.0 / 7.0
	// var r float64
	// var luas float64

	// fmt.Println("input r")
	// fmt.Scan(&r)

	// luas = (4 * pi) * (r * r)

	// fmt.Println(luas)

	//soalNomor2

	// var d1, d3 int
	// var d2, d4 int

	// fmt.Println("input 2 digits")
	// fmt.Scan(&d1, &d3)

	// d2 = d1 / 1
	// d4 = d3 / 1

	// fmt.Println(d1, d2, d3, d4)

	// // soalNomor3

	var in string
	var x byte
	var out bool

	fmt.Println("input data")
	fmt.Scan(&in)

	x = in[0]

	out = x >= 65 && x <= 90

	fmt.Println(out)

	// //soalNomor4

	// var E0, E1, c int
	// var cc int

	// fmt.Scan(&E0, &c, &E1)

	// cc = (E0 - E1) / c

	// fmt.Println(cc)

}
