package main

import "fmt"





func main() {

	var palin int
	var found bool

	fmt.Print("enter a 6 digit number greater than 101101:")
	fmt.Scanln(&palin)

	palin= InitPalin(palin)

	for !found{
		fmt.Println(palin)
		found = TestPalin(palin)

		palin = NextPalin(palin)


	}
}

func InitPalin(pal int)int{



	hundK := pal/100000
	tenK := (pal/10000)%10
	onek := (pal/1000)%10

	num1 := hundK+(tenK*10)+(onek*100)
	num2 := pal%1000

	if num1 >num2{
		pal = (pal/1000-1)*1000

		hundK = pal/100000
		tenK = (pal/10000)%10
		onek = (pal/1000)%10

		num1 = hundK+(tenK*10)+(onek*100)
	}
	pal = (pal/1000)*1000+num1


	return pal
}

func NextPalin(pal int)int{


	pal = (pal/1000-1)*1000

	hundK := pal/100000
	tenK := (pal/10000)%10
	onek := (pal/1000)%10

	num1 := hundK+(tenK*10)+(onek*100)

	pal = (pal/1000)*1000+num1
	return pal

}

func TestPalin(pal int) bool{

	var isFound bool

	for i:=999 ;i>=100;i--{
		if pal % i == 0{
			fmt.Println(pal%i)
			num :=pal/i
			fmt.Println(i*num, i,num)
			if num>=100&&num<=999{
				isFound = true
				break
			}
		}

	}
	return isFound
}