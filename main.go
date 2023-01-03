package main

import "fmt"

func main() {
	var number, NUM int
	var reverse int = 0

	fmt.Print("Enter any integer : ")
	fmt.Scan(&number)

	NUM = number
	for {

		reverse = reverse*10 + number%10
		number /= 10

		if number == 0 {
			break
			// Break Statement it comes out of the loop
		}
	}

	if NUM == reverse {
		fmt.Printf("%d is a Palindrome \n", NUM)
	} else {
		fmt.Printf("%d is not a Palindrome \n", NUM)
	}

}
