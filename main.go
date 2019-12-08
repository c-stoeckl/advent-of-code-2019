package main

import (
	"log"
	"os"
	"strconv"

	AOC "github.com/c-stoeckl/advent-of-code-2019/src"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("Please provide the number of the day for which you want the solution")
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	AOC.Solve(day)
}
