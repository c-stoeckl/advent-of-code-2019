package adventofcode

import (
	"bufio"
	"os"
	"strconv"
)

// ParseInput reads a whole file into memory
// and returns a slice of its lines.
func parseInput(day int) ([]string, error) {
	file, err := os.Open("adventofcode/inputs/" + strconv.Itoa(day) + ".txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
