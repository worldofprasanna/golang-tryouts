package main

import (
	"fmt"
	"os"
)

func main() {
	if (len(os.Args) != 2) {
		fmt.Println("Specify argument")
		os.Exit(1)
	}
	fmt.Println("This is the argument: ", os.Args)
}
