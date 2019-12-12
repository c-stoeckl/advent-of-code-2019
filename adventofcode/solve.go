package adventofcode

import "fmt"

// Solve calls the correct solve function for the given day
func Solve(day int) {
	printGreeting(day)

	switch day {
	case 1:
		day1()
	case 2:
		day2()
	default:
		fmt.Println("No solution for the provided day, yet.")
	}
}

func printGreeting(day int) {
	var divider string
	addition := "*****"
	greeting := fmt.Sprintf("%s Day %d %s", addition, day, addition)

	for i := 0; i < len([]rune(greeting)); i++ {
		divider += "*"
	}

	fmt.Println(divider)
	fmt.Println(greeting)
	fmt.Println(divider)
}
