package main

import (
	"bufio"
	"fmt"
	"os"
)

/* path of file containing puzzle input*/
const inputFile string = "input2.txt"

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

func isLevenshteinDistEqualsOne(firstString string, secondString string) (bool, int) {
	mismatchCount := 0
	mismatchIndex := -1
	i := 0
	// fmt.Println("first string: ", firstString, "\nSecond string: ", secondString)
	for i < len(firstString) {
		if firstString[i] != secondString[i] {
			// fmt.Println("found mismatch")
			mismatchIndex = i
			mismatchCount++
		}
		if mismatchCount > 1 {
			return false, mismatchIndex
		}
		i++
	}
	// fmt.Println("Found our box. returning:", mismatchIndex)
	return true, mismatchIndex
}

func findCorrectBoxID(lines []string) (string, int) {
	i := 0
	for i < len(lines) {
		j := i + 1
		for j < len(lines) {
			firstString := lines[i]
			secondString := lines[j]
			// fmt.Println(i, lines[i])
			// fmt.Println(j, lines[j])
			ret, index := isLevenshteinDistEqualsOne(firstString, secondString)
			if ret {
				return firstString, index
			}
			j++
		}
		i++
	}
	return "", -1
}

func countRepeats(line string) (bool, bool) {
	characterFrequency := make(map[byte]int, 0)
	twoCount := 0
	threeCount := 0

	i := 0
	for i < len(line) { /* loop till the end of the line containing characters*/
		c := line[i]

		characterFrequency[c]++
		if characterFrequency[c] == 2 { /* Found our character repeating 2 times */
			twoCount++ /* increment two counter */

		} else if characterFrequency[c] == 3 { /* Found character repeating 3 times */
			threeCount++ /* increment three counter */
			twoCount--   /* decrement two counter for the same character */

		} else if characterFrequency[c] == 4 { /* shit it repeated 4 times, decrement three count */
			threeCount--

		}
		i++
	}
	return twoCount != 0, threeCount != 0 /* Convert int to bool and return. false if 0, true for anything else*/
}

func findChecksum(lines []string) int {
	twoCount := 0
	threeCount := 0

	i := 0
	for i < len(lines) {
		l := lines[i]
		if len(l) == 0 {
			continue
		}

		containsTwo, containsThree := countRepeats(lines[i])
		if containsTwo {
			twoCount++
		}
		if containsThree {
			threeCount++
		}
		i++
	}

	fmt.Println("total two counts: ", twoCount)
	fmt.Println("total three counts: ", threeCount)
	checksum := twoCount * threeCount
	return checksum
}

func main() {
	lines, err := readLines(inputFile)
	check(err)
	fmt.Println("Checksum: ", findChecksum(lines))
	id, index := findCorrectBoxID(lines)
	fmt.Println("found the correct Box ID:", id, "with wrong index:", index)
}
