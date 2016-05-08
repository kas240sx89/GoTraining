package main

import "fmt"

func main() {

	var str string
	var numDig int =10
	var consec int = 5
	var largestVal int
	var arr []int
	fmt.Print("please enter a 10 digit positive number:")
	fmt.Scanln(&str)

	arr = StrToIntArr(str, numDig)


	largestVal = LargestConsecProduct(arr,numDig,consec)

	fmt.Println("largest Consec Product is:",largestVal)
}


func StrToIntArr(str string,num int)[]int{

	runes := []rune(str)
	var temp []int

	for i:= 0;i<num;i++{
	 	runes[i] = runes[i]-48
		temp = append(temp,int(runes[i]))
	}

	return temp
}

func LargestConsecProduct(arr []int, numD int, numC int)int{

	var temp int
	var ans int

	for i := 0; i <=(numD- numC);i++{

		temp = 1

		for j:= 0;j<numC;j++{
			temp = temp * arr[i+j]
		}

		if temp> ans{
			ans = temp
		}

	}

	return ans

}