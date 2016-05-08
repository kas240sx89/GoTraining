package main

import "fmt"

func main() {

	var num int
	primes := []int64{}



	fmt.Println("------Largest Prime Factor------")
	fmt.Print("Please enter a positive number: ")
	fmt.Scanln(&num)

	primes =GetPrimes(num,primes)


		fmt.Println("The",num, "is:",primes[num])

}



func IsPrimeNum(n int64,primes[]int64) bool{

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

func GetPrimes(num int, primes []int64)[]int64 {
	var temp int64 =2

	for i:= 0; i<=(num);{

		isPrime :=IsPrimeNum(temp,primes)

		if isPrime==true{
			primes = append(primes,temp)
			//fmt.Println(i)
			i++
		}

		temp++

	}
	return primes
}
