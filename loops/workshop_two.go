package loops

import "fmt"

//Goal: Prediction game. Make prediction untill find the number.
func Workshop_two() {
	var secretNumber int = 55
	fmt.Println("I kept a number. Can you find it?")
	for true {
		var guess int
		fmt.Scanln(&guess)

		if guess == secretNumber {
			fmt.Println("You find it!")
			break
		}

		if guess < secretNumber {
			fmt.Println("Up!")
		} else {
			fmt.Println("Down!")
		}
	}
}
