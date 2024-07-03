package main

import (
	"fmt"
)

var a int
var b int = 2
var c = 3

const (
	A int = 1
	B     = 3.14
	C     = "Hi!"
)

func main() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred
	a = 1

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println("Hello World!")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
