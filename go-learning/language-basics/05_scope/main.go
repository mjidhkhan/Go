package main

import "fmt"

//var  number := 42 // con not be assigned

var num = 42 // package scope

var num1 int // package scope

//number2 = 42 // will nor work

func main() {
	num1 = 2 // initialized here
	num2 := 50
	fmt.Println(num, num1, num2)
	foo()
}

func foo() {
	fmt.Println("Called from  foo() num and num1 ", num, num1)

	// no access to num2
	// this does not compile
	fmt.Println(num2)
}
