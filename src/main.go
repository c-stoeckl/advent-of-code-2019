package adventofcode

import "fmt"

// Solve calls the correct solve function for the given day
func Solve(day int) {
	switch day {
	case 1:
		day1()
	default:
		fmt.Println("No solution for the provided day, yet.")
	}
}
