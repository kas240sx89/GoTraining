package main

import "fmt"

func main(){

	var num int
	var sum int


	fmt.Print("Please enter a number:")
	fmt.Scanln(&num)

	for i := 0; i< num;i++{

		if i % 3 == 0 || i % 5 ==0{
			sum+=i
		}
	}

	fmt.Println("The sum of the multiples of 3 or 5, between 0 and" , num , "is: ",sum)

}