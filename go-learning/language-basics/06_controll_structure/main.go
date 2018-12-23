package main

import "fmt"

func main() {

	// loop

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	for x := 1; x <= 20; x++ {
		// switch
		var y = x % 2
		var e int
		if y == 0 {
			e = 0
		} else {
			e = 1
		}

		switch e {
		case 0:
			fmt.Println("even")
		case 1:
			fmt.Println("odd")
		default:
			fmt.Println("hard to figure it out")
		}
	}

}
