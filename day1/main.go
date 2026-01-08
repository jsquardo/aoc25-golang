package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Open file input
	file, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Initialize state
	position := 50 // starting number
	hits := 0      // number of times it lands on 0

	for scanner.Scan() {
		line := scanner.Text()

		direction := line[0]
		// Distance is everything after the first character
		distanceStr := line[1:]
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			panic(err)
		}

		hits += countZeroHits(position, direction, distance)

		// Move the "dial"
		if direction == 'R' {
			position += distance
		} else if direction == 'L' {
			position -= distance
		}

		// Wrap around dial 0-99
		position = position % 100
		if position < 0 {
			position += 100
		}

	}

	// Check scanner for errors
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Answer:", hits)
}

// Function to help part 2
func countZeroHits(pos int, dir byte, dist int) int {
	var first int

	if dir == 'R' {
		first = (100 - pos) % 100
	} else { // L
		first = pos % 100
	}

	if first == 0 {
		first = 100
	}

	if dist < first {
		return 0
	}

	return 1 + (dist-first)/100
}
