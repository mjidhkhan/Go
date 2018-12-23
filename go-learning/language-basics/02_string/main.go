package main

import "fmt"

func main() {
	fmt.Println("Hello, World"[0])       // print 72 H
	fmt.Println(" "[0])                  // print 32 space
	fmt.Println(len("Hello, World"))     // print 12 not 11 as space is also a character
	fmt.Println(len("Hello," + "World")) // print 11 no space
}
