package main

import "fmt"

func main() {
	var x int
	var y int
	var z int
	fmt.Println("enter 1st number:")
	fmt.Scanln(&x)
	fmt.Println("enter 2nd number:")
	fmt.Scanln(&y)
	fmt.Println("enter 3rd number:")
	fmt.Scanln(&z)
	sum := x + y + z
	avg := float64(sum) / 3
	fmt.Println("Среднее арифметическое:", avg)
}
