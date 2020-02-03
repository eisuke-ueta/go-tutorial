package main

import (
	"fmt"
	"os/user"
	"time"
)

const Pi = 3.14
const (
	Username = "test_user"
	Password = "test_pass"
)

func main() {
	// Hello World
	fmt.Println("Hello world", time.Now())
	fmt.Println(user.Current())

	// Variables
	var i int = 1
	var f64 float64 = 1.2
	var s string = "test"
	var t bool = true
	fmt.Println(i, f64, s, t)

	xi := 1
	xf64 := 1.2
	xs := "test"
	xt := true
	fmt.Println(xi, xf64, xs, xt)
	fmt.Printf("%T\n", xf64)
}
