package main

import (
	"fmt"
)

func main() {

	var lower int =1
	var upper int =1
	var temp int
	var maxVal int
	var sum int

	fmt.Print("Please enter the maximum value of the terms int the fibonacci sequence: " )
	fmt.Scanln(&maxVal)

	fmt.Print(lower	)
	for  upper<maxVal{
		fmt.Print("," , upper )

		if upper % 2 == 0{
			sum+= upper
		}

		temp = lower
		lower = upper
		upper = temp + lower

	}

	fmt.Println("\nThe sum of the even terms below",maxVal,"is:",sum)



}
