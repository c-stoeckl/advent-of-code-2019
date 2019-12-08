package main

import (
	"log"
	"os"

	AOC "github.com/c-stoeckl/advent-of-code-2019/src"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("Please provide the number of the day for which you want the solution")
	}

	AOC.Solve(1)
}
