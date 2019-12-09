package adventofcode

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// ParseInput reads a whole file into memory
// and returns a slice of its lines.
func parseInput(day int) []string {
	file, err := os.Open("adventofcode/inputs/" + strconv.Itoa(day) + ".txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Invalid input: %s", err)
	}

	return lines
}

func printSolution(solutions ...interface{}) {
	for i, solution := range solutions {
		fmt.Printf("Part %d: %v\n", i+1, solution)
	}
}
