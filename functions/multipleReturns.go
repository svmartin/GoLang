package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	theSquareRoot, err := squareRoot(9) // result, err = nil if no error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(theSquareRoot)
}

func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("can't do square root of negative number")
	}
	return math.Sqrt(x), nil // if no error, returns square root and mil
				// if error, returns error message
}
