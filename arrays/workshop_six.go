package arrays

import "fmt"

//GOAL: Create a 2 dimensional array and print that with correct alignment.
func Workshop_six() {
	var arr [2][3]int
	arr[0][0] = 1
	arr[0][1] = 3
	arr[0][2] = 5
	arr[1][0] = 2
	arr[1][1] = 4
	arr[1][2] = 6

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
