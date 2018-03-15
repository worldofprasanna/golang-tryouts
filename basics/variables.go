package main

import "fmt"

func main() {
	var i int = 10
	j := 20

	fmt.Println("Values of i & j ", i , j)

	a, b, i := 10, 20, 30
	fmt.Println("Values of i, j, a, b ", i , j, a, b)

	//i := 40 Compilation error
}
