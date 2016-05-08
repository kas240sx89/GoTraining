package main

import "fmt"

func main() {
	var num int
	var found bool
	var ans int
	primes := []int{1}
	var pivit int = 1

	fmt.Print("Please enter a positive number up to 40:")
	fmt.Scan(&num)

	primes = GetPrimes(num,primes)
	//fmt.Println(primes[0]...)

	for i := 0;i <len(primes);i++{

		pivit = pivit * primes[i]
	}



	if num>2 {

		for !found{

			ans += pivit

			found = TestAns(num,ans)

		}


	}

	fmt.Println(ans)

}

func TestAns(num int,ans int) bool{

	var found bool

	for i := num; i >= 0; i-- {

		if i ==0{
			found=true
		}else if (ans % i) != 0 {
			break
		}

	}
	return found
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