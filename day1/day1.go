package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/* path of file containing puzzle input*/
const inputFile string = "input1.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
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

func calculateFinalFrequency(lines []string, initialFrequency int) int {
	currentFrequency := initialFrequency
	for i := range lines {
		frequencyDrift, err := strconv.Atoi(lines[i])
		check(err)
		currentFrequency += frequencyDrift
	}

	return currentFrequency
}

func findFirstRepeatingFrequency(lines []string, initialFrequency int) (int, bool) {
	/* make a map of integers, integer key-value. Will store freq count of sums in each steps */
	history := make(map[int]int, 0)

	currentFrequency := initialFrequency
	found := false

	i := 0
	for i < len(lines) { /* loop till the end of the frequency change list*/
		frequencyDrift, err := strconv.Atoi(lines[i])
		check(err)
		if len(lines[i]) == 0 {
			continue
		}

		currentFrequency += frequencyDrift

		history[currentFrequency]++
		if history[currentFrequency] == 2 { /* Found our first repeating sum! break out and return */
			found = true
			break
		}
		if i == (len(lines) - 1) { /* restart the loop until we get our repeating sum */
			i = 0
		} else {
			i++
		}
	}

	return currentFrequency, found
}

func main() {
	lines, err := readLines(inputFile)
	check(err)
	fmt.Println("Resulting frequency after all fluctuations: ", calculateFinalFrequency(lines, 0))
	firstRepeatingFrequency, found := findFirstRepeatingFrequency(lines, 0)
	if found {
		fmt.Println("Correct frequency after callibration: ", firstRepeatingFrequency)
	} else {
		fmt.Println("No repeating numbers")
	}
}
