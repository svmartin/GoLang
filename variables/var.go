package main

import "fmt"

var guac = "avocado"

func main() {
	var a int
	a = 1
	var b, c int
	b = 2
	c = 3
	var d = 5
	e := 6
	f, g := 8, 9
	var wholeNumber int = 11
	var fractionalNumber float64 = 2.75
	var wholeNumber2  int = int(fractionalNumber)
	fmt.Println(a, b, c, d, e, f, g)
	fmt.Println(wholeNumber)
	fmt.Println(fractionalNumber)
	fmt.Println(wholeNumber2)
	fmt.Println(float64(wholeNumber) + fractionalNumber)
	fmt.Println(float64(wholeNumber2) < fractionalNumber)
	fmt.Println(guac)

}