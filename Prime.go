package main

import (
	"fmt"
)
func main() {
	var number int
	fmt.Print("Enter Number To check Prime or not:")
	fmt.Scanln(&number)
	Prime(number)
}

func Prime(number int) {
	isPrime := true
	if number == 0 || number == 1 {
		
	} else {
		for i := 2; i < number/2; i++ {
			if number%i == 0 {
				fmt.Printf(" %d is not a  prime number\n", number)
				isPrime = false
				break
			}
		}
		if isPrime == true {
			fmt.Printf(" %d is a prime number\n", number)
		}
	}
}
