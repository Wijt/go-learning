package conditionals

import "fmt"

//Goal: Declare 3 variables and find biggest one.
func Workshop_one() {
	var one int = 17
	var two int = 25
	var three int = 22

	var biggest = one

	if biggest < two {
		biggest = two
	}

	if biggest < three {
		biggest = three
	}

	fmt.Println("Biggest number is", biggest)
}
