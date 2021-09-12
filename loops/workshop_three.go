package loops

import "fmt"

//Goal: Prediction game. Make prediction until find the number then give user score.
//1-3 Perfect!, 4-6 Nice! 6> Good enough!
func Workshop_three() {
	var secretNumber int = 55
	var guessCount int = 0

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
		guessCount++
	}

	if guessCount <= 3 {
		fmt.Println("Perfect!")
	} else if guessCount <= 6 {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Good Enough!")
	}
}
