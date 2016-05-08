package main

import "fmt"

func main() {


	var sumOfSqs int
	var SqOfSums int
	var num int


	fmt.Print("Please enter the a integer less than 50: " )
	fmt.Scanln(&num)

	sumOfSqs = GetSumOfSquares(num)
	SqOfSums = GetSqOfSums(num)


	fmt.Println((SqOfSums-sumOfSqs))


}

func GetSumOfSquares(num int)int{

	var ans int

	for i:=1;i<=num;i++{
		ans+= (i*i)

	}

	return ans
}

func GetSqOfSums(num int)int{

	var ans int

	for i:=1;i<=num;i++{
		ans+= i

	}

	return (ans* ans)
}