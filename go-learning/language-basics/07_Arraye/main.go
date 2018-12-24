package main

import "fmt"

func main() {
	//arrays
	var x [5]int
	x[4] = 100
	fmt.Println(x)    // will print [0 0 0 0 100]
	fmt.Println(x[4]) // will print 100

	// sum array elements
	var y [5]float64
	y[4] = 100
	y[3] = 99
	y[2] = 98
	y[1] = 97
	y[0] = 96
	var total float64
	for i := 0; i < len(y); i++ {
		fmt.Println(i)
		total += y[i]

	}
	fmt.Println("Summ of array : ", total)
	// average
	var avg = total / float64(len(y)) // type casting int to float64
	fmt.Println("Average : ", avg)

	total = 0
	// another way off adding int array values
	for _, value := range y {
		total += value
	}
	fmt.Println(total)

	// short hand way to create arrays
	z := [10]float64{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(z)

	//SLICE (segment of an array)

	w := make([]float64, 5, 10)
	fmt.Println(w[1:5])
	fmt.Println(w)
	slice1 := z[0:5]
	slice2 := append(slice1, 120, 130)
	fmt.Println(slice1, slice2)
}
