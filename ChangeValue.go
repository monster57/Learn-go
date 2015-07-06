package main

import "fmt"

type Add struct{
	number1 int
	number2 int
}
func main() {
	num := Add{2,211}
	num2 := num.number1+num.number2
	fmt.Println("this is the add " , num	)
	fmt.Println("this is the value" , num2)
}
