package loops

import "fmt"

//Get a number from user then find it prime or not.
func Workshop_four() {
	var number int
	fmt.Print("Number: ")
	fmt.Scan(&number)
	var prime bool = true

	for i := 2; i < number; i++ {
		if number%i == 0 {
			prime = false
			break
		}
	}

	if prime {
		fmt.Println(number, "is a prime number.")
	} else {
		fmt.Println(number, "is not a prime number.")
	}
}
