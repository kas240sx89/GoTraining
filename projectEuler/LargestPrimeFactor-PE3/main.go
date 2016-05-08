package main

import "fmt"

func main() {

	var num int
	primes := []int{}

	primes =GetPrimes(num,primes)

	fmt.Println("------Largest Prime Factor------")
	fmt.Print("Please enter a positive number: ")
	fmt.Scanln(&num)
	//fmt.Print(1)

	primes =GetPrimes(num,primes)


	PrintLargestPrimeFactor(num,primes)

}


func IsPrimeNum(n int,primes[]int) bool{

	 isPrime :=true

	if n >3 {
		for p := 1; p < len(primes); p++ {
			pri := primes[p]

			if n % pri == 0 {
				isPrime = false
				break
			}
		}
	}
	return isPrime
}

func GetPrimes(num int, primes []int)[]int {

	for i:= 2; i<=num;i++{

		isPrime :=IsPrimeNum(i,primes)

		if isPrime==true{
			primes = append(primes,i)
			//fmt.Print("," ,i)
		}

	}
 return primes
}

func PrintLargestPrimeFactor(num int,primes []int){

	for i :=len(primes)-1;i>=0;i--{

		pri := primes[i]

		if num% pri ==0{
			fmt.Println("The Largest Prime Factor of",num,"is:",pri)
			break
		}
	}

}
