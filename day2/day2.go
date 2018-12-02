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
}
