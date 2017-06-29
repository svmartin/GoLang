package main

import "fmt"


func main() {
	myFunction()
	fmt.Println(square(3))
	fmt.Println(add(3, 4))
	fmt.Println(subtract(10, 22))
}

func myFunction() {
	fmt.Println("Running myFunction y'all")
}

func square(number int) int {
	return number * number
}

func add(a, b float64) (sum float64){
	return a + b
}

func subtract(a, b float64) (difference float64){
	difference = a - b
	return // bare return
}