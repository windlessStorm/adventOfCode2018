package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync"
)

/* path of file containing puzzle input*/
const inputFile string = "input5.txt"

type stack struct {
	lock sync.Mutex
	data []byte
}

func NewStack() *stack {
	return &stack{sync.Mutex{}, make([]byte, 0)}
}

func (s *stack) Push(v byte) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data = append(s.data, v)
}

func (s *stack) Pop() (byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.data)
	if l == 0 {
		return 0, errors.New("Empty stack")
	}
	res := s.data[l-1]
	s.data[l-1] = 0
	s.data = s.data[:l-1]
	return res, nil
}

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

// swapCase upper to lower or lower to upper
func swapCase(b byte) byte {
	switch {
	case 97 <= b && b <= 122:
		return b - 32
	case 65 <= b && b <= 90:
		return b + 32
	default:
		return b
	}
}

func lowerCase(b byte) byte {
	switch {
	case 65 <= b && b <= 90:
		return b + 32
	default:
		return b
	}
}

func reactPolymer(s *stack, input string, ignoreLetter byte) {
	fmt.Println("Will ignore letter:", string(ignoreLetter))
	for i := 0; i < len(input); i++ {
		if lowerCase(input[i]) == ignoreLetter {
			continue
		}
		if len(s.data) == 0 {
			s.Push(input[i])
			continue
		}
		if swapCase(input[i]) == s.data[len(s.data)-1] {
			s.Pop()
			continue
		}
		s.Push(input[i])
	}
}

func findBestShrink(input string) int {
	var i byte
	bestShrinkSize := len(input)
	for i = 97; i <= 122; i++ {
		s := NewStack()
		reactPolymer(s, input, i)
		if len(s.data) < bestShrinkSize {
			bestShrinkSize = len(s.data)
		}
	}
	return bestShrinkSize
}

func main() {
	plans, err := readLines(inputFile)
	check(err)
	input := plans[0]
	s := NewStack()
	reactPolymer(s, input, 0)
	fmt.Println("Part1:", len(s.data))
	bestShrinkSize := findBestShrink(input)
	fmt.Println("Part2:", bestShrinkSize)

}
