package main

import "fmt"

func main() {

	var number int
	var reverse int = 0

	fmt.Print("Enter any integer : ")
	fmt.Scan(&number)

	for {

		reverse = reverse*10 + number%10
		number /= 10

		if number == 0 {
			break // to come out of the loop

		}
	}

	if number == reverse {
		fmt.Printf("%d is a Palindrome \n", number)
	} else {
		fmt.Printf("%d is not a Palindrome \n", number)
	}

}
