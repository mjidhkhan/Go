package main

import "fmt"

func main() {

	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Println()

	var x = "Hello, world"
	fmt.Println(x)

	var y string
	y = "first"
	println(y)
	y = y + " Second"
	println(y)
	y = y + " third"

	println(y)

	// not good naming practice such as a, b, c, d or x, y ,z
	dl := "dolly"
	fmt.Println("My cat name is ", dl)
	// better way to write descriptive name
	name := "hachi"
	fmt.Println("My dog name is ", name)
	// or
	dogName := "pat"
	fmt.Println("My dog name is ", dogName)

	// define multipls variables
	var (
		v1 = 3
		v2 = "test"
		v3 = 10
	)

	fmt.Println("Variables var( v1 v2 v3 )", v1, v2, v3)

	// about const

	const greet = "Hello World"
	fmt.Println(greet)

	//greet = "another greet " // will generat error cnnot assing to greet, which is constant

}
