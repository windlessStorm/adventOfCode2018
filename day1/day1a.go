package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/* path of file containing puzzle input*/
const inputFile string = "input1a.txt"

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
	var currentFrequency = initialFrequency
	for i := range lines {
		frequencyDrift, err := strconv.Atoi(lines[i])
		check(err)
		currentFrequency += frequencyDrift
	}
	return currentFrequency
}

func main() {
	lines, err := readLines(inputFile)
	check(err)
	fmt.Println(calculateFinalFrequency(lines, 0))
}
