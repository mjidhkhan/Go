/*
Write a program that prints the numbers from 1 to 100,
but for multiples of three, print “Fizz” instead of
the number, and for the multiples of five, print “Buzz.”
For numbers that are multiples of both three and five,
print “FizzBuzz.”
*/

package main

import "fmt"

func main() {

	/*
		Write a program that prints out all the numbers between 1 and 100
		that are evenly divisible by 3 (i.e., 3, 6, 9, etc.).
	*/

	for i := 1; i <= 100; i++ {

		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	/*
		Write a program that prints the numbers from 1 to 100,
		but for multiples of three, print “Fizz” instead of
		the number, and for the multiples of five, print “Buzz.”
		For numbers that are multiples of both three and five,
		print “FizzBuzz.”
	*/

	for x := 1; x <= 100; x++ {
		if x%3 == 0 && x%5 == 0 {
			println("FizzBuzz")
		} else if x%5 == 0 {
			println("Buzz")
		} else if x%3 == 0 {
			println("Fizz")
		} else {
			println(x)
		}
	}
}
