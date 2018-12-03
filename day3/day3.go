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

/* 
Single pass solution:

This draws the plans on our fabric, also keeps track of overlapping cell counts, and filter out the 
best plan ID, which is non-overlapping rectangle. 

All rectangles/plans, which we draw participates in best plan competition, we keep adding and deleting 
potentially best plan to a list. At the end of this function execution, 
just the BEST of the BEST only non-overlapping plan will remain in that list- our winner
*/
func drawPlans(plans []string, fabric [][]int) ([][]int, int, map[int]bool) {
	i := 0
	overlapArea := 0
	hasOverlap := false
	bestFabricID := map[int]bool{} /* Best plan competition, keep adding and deleting potential best plans */
	for i < len(plans) {
		id, marginLeft, marginTop, width, hieght:= parsePlan(plans[i])
		for j:= marginLeft; j < marginLeft + width; j++ {
			for k:= marginTop; k < marginTop + hieght; k++ {
				if fabric[j][k] == 0 { /* Cell not part of any plan till now, mark it as part of current plan*/
					fabric[j][k] = id
				} else if	fabric[j][k] == -1 { /* Already a overlapping cell*/
						hasOverlap = true
						continue
					} else { /* Already marked by some other plan, make this -1 to show now it is overlapping*/
						hasOverlap = true

						/* if we overlap on some ID which was contender for best plan, then it is not best plan for sure,
						Delete it from best plan list */
						if bestFabricID[fabric[j][k]] { /* check if id is in best plan competition*/
							delete(bestFabricID, fabric[j][k])
						}
						fabric[j][k] = -1
						overlapArea++
					}
			}
		}
		if !hasOverlap{
			bestFabricID[id] = true
			// fmt.Println("Best fabric set now have:", bestFabricID)
		}
		hasOverlap = false /* Reset the overlap flag*/
		i++
	}
	// fmt.Println("All done and dusted we have a winner over here guys and it isss:", bestFabricID )
	return fabric, overlapArea, bestFabricID
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
	fabric, overlapArea, bestFabricID := drawPlans(plans, fabric)
	// fmt.Println(fabric)
	fmt.Println("Overlaping area:", overlapArea)
	fmt.Println("Best fabric piece without any overlap is:", bestFabricID)
}
