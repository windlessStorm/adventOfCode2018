package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

/* path of file containing puzzle input*/
const inputFile string = "input3.txt"

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

func parsePlan(plan string) (int, int, int, int, int) {
	planList := strings.Split(plan, " ")
	id, err := strconv.Atoi(planList[0][1:])
	check(err)
	margin := strings.Split(planList[2], ",")
	marginLeft, err := strconv.Atoi(margin[0])
	check(err)
	marginTop, err := strconv.Atoi(margin[1][:(len(margin[1])-1)])
	check(err)
	dimension := strings.Split(planList[3], "x")
	width, err := strconv.Atoi(dimension[0])
	check(err)
	hieght, err := strconv.Atoi(dimension[1])
	check(err)
	return id, marginLeft, marginTop, width, hieght
}

func drawPlans(plans []string, fabric [][]int) ([][]int, int) {
	i := 0
	overlapArea := 0
	for i < len(plans) {
		id, marginLeft, marginTop, width, hieght:= parsePlan(plans[i])
		for j:= marginLeft; j < marginLeft + width; j++ {
			for k:= marginTop; k < marginTop + hieght; k++ {
				if fabric[j][k] == 0 {
					fabric[j][k] = id
				} else if	fabric[j][k] == -1 {
						continue
					} else {
					fabric[j][k] = -1
					overlapArea++
				}
			}
 		}
		i++
	}
	return fabric, overlapArea
}

func createFabric(totalRows int, totalColumns int) [][]int {
	fabric := make([][]int, totalRows)
  for i := 0; i < totalRows; i++ {
		fabric[i] = make([]int, totalColumns)
	}
	return fabric
}


func main() {
	plans, err := readLines(inputFile)
	check(err)
	fabric := createFabric(1000, 1000)
	fabric, overlapArea := drawPlans(plans, fabric)
	// fmt.Println(fabric)
	fmt.Println("Overlaping area:", overlapArea)
}
