package main

import "fmt"

func main() {

	fmt.Printf("%d - %b \n", 20, 20)
	fmt.Printf("%d - %b - %o - %#x \n", 44, 44, 44, 44)


	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %o \t %#x %q \n", i, i, i, i, i)
	}



}
