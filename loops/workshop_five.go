package loops

import "fmt"

//Goal: Write a program that finds 2 given variables are amicable numbers.
func Workshop_five() {
	var numberOne int = 1184
	var numberTwo int = 1210

	var numberOneSum int = 0
	for i := 1; i <= numberOne/2; i++ {
		if numberOne%i == 0 {
			numberOneSum += i
		}
	}

	var numberTwoSum int = 0
	for i := 1; i <= numberTwo/2; i++ {
		if numberTwo%i == 0 {
			numberTwoSum += i
		}
	}

	if numberOne == numberTwoSum && numberTwo == numberOneSum {
		fmt.Println(numberOne, "and", numberTwo, "is an amicable number.")
	} else {
		fmt.Println(numberOne, "and", numberTwo, "is not an amicable number.")
	}
}
