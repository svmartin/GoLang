package main

import (
	"fmt"
	"math"
	"github.com/svmartin/welcome"
	"github.com/golang/example/stringutil"
	"reflect"
	"net"
	"time"
)

// go get github.com/golang/example/stringutil

var myNumber = 1.23

func main() {
	roundedUp := math.Ceil(myNumber)
	roundedDown := math.Floor(myNumber)
	fmt.Println(roundedUp, roundedDown)
	fmt.Println(welcome.English)
	fmt.Println(welcome.Spanish)
	fmt.Println(stringutil.Reverse(welcome.Spanish))
	//fmt.Println("hello" + 787)
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(1.5))
	fmt.Println(reflect.TypeOf("hello" ))
	fmt.Println(reflect.TypeOf(false))
	fmt.Println(reflect.TypeOf(net.IPv4(127, 0, 0, 1)))
	fmt.Println(reflect.TypeOf(time.Now()))
}