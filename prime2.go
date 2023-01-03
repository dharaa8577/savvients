package main

import "fmt"

func main() {

	var primNum, primcount int
	primcount = 0

	fmt.Print("Enter the Number to find the Prime Numbers = ")
	fmt.Scanln(&primNum)

	for i := 2; i < primNum/2; i++ {
		if primNum%i == 0 {
			primcount++
			break
		}
	}

	if primcount == 0 && primNum != 1 {
		fmt.Println(primNum, " is a Prime Number")
	} else {
		fmt.Println(primNum, " is Not a Prime Number")
	}
}
