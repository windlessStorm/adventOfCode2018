package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/* path of file containing puzzle input*/
const inputFile string = "input4.txt"

type recordEntry struct {
	timeStamp time.Time
	event     string
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

func parseLog(lines []string) []recordEntry {
	var log []recordEntry

	i := 0
	for i < len(lines) {
		// fmt.Println(lines[i])
		var record recordEntry
		timeStamp, err := time.Parse("2006-01-02 15:04", lines[i][1:17])
		record.timeStamp = timeStamp
		check(err)
		record.event = lines[i][19:]
		// fmt.Println(record.timeStamp.String(), record.event)
		log = append(log, record)
		i++
	}
	sort.Slice(log, func(i, j int) bool { return log[i].timeStamp.Before(log[j].timeStamp) })
	return log
}

func computeSleepData(log []recordEntry) (map[string]int, map[string][60]int) {
	sleepData := map[string]int{}
	sleepPattern := map[string][60]int{}
	var id string
	var sleepTime time.Time

	i := 0
	for i < len(log) {
		// fmt.Println(log[i].timeStamp.String(), log[i].event)
		eventType := log[i].event[:5]
		if eventType == "Guard" {
			id = strings.Split(log[i].event, " ")[1]
			// fmt.Println(id)
		}
		if eventType == "falls" {
			sleepTime = log[i].timeStamp
			// fmt.Println(sleepTime)
		}
		if eventType == "wakes" {
			wakeTime := log[i].timeStamp
			// timeSlept := int(wakeTime.Sub(sleepTime).Minutes())
			timeArr := sleepPattern[id]
			for k := sleepTime.Minute(); k < wakeTime.Minute(); k++ {
				timeArr[k]++
			}
			sleepPattern[id] = timeArr
			timeSlept := wakeTime.Minute() - sleepTime.Minute()
			sleepData[id] += timeSlept
		}
		// fmt.Println(id)

		i++
	}
	return sleepData, sleepPattern
}

func getPart1(sleepData map[string]int, sleepPattern map[string][60]int) int {
	var max int
	var maxID string
	for k := range sleepData {
		if sleepData[k] > max {
			max = sleepData[k]
			maxID = k
		}
	}
	max = -1
	sleepyTime := -1
	for k := range sleepPattern[maxID] {
		if sleepPattern[maxID][k] > max {
			max = sleepPattern[maxID][k]
			sleepyTime = k

		}
	}
	id, err := strconv.Atoi(maxID[1:])
	check(err)

	return id * sleepyTime
}

func getPart2(sleepPattern map[string][60]int) int {
	sleepiestTime := -1
	sleepiestDuration := -1
	sleepiestID := -1
	for k := range sleepPattern {
		for j := range sleepPattern[k] {
			if sleepPattern[k][j] > sleepiestDuration {
				sleepiestTime = j
				sleepiestDuration = sleepPattern[k][j]
				var err error
				sleepiestID, err = strconv.Atoi(k[1:])
				check(err)
			}
		}
	}
	return sleepiestID * sleepiestTime
}

func main() {
	lines, err := readLines(inputFile)
	check(err)

	log := parseLog(lines)
	sleepData, sleepPattern := computeSleepData(log)

	fmt.Println(getPart1(sleepData, sleepPattern))
	fmt.Println(getPart2(sleepPattern))

}
