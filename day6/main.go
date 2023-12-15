package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	duration int
	distance int
}

func main() {
	file, err := os.Open("part1.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	races, race := parseRaces(&lines)

	sum := 0
	for _, v := range races {
		simulation := simulateRace(v)
		if sum == 0 {
			sum = len(simulation)
		} else {
			sum *= len(simulation)
		}
	}

	single_sum := len(simulateRace(race))

	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %d\n", single_sum)
}

func parseNumbers(chars []string) []int {
	nums := make([]int, 0)
	num := ""
	for i := 0; i < len(chars); i++ {
		if chars[i] == " " {
			if num != "" {
				n, err := strconv.Atoi(num)
				if err != nil {
					continue
				}

				nums = append(nums, n)
				num = ""
			}
			continue
		}

		n, err := strconv.Atoi(chars[i])
		if err != nil {
			continue
		}
		num += strconv.Itoa(n)
	}

	if num != "" {
		n, _ := strconv.Atoi(num)
		nums = append(nums, n)
	}

	return nums
}

func parseRaces(lines *[]string) ([]Race, Race) {
	time_string := (*lines)[0]
	distance_string := (*lines)[1]

	time_split := strings.Split(strings.Split(time_string, "Time: ")[1], "")
	distance_split := strings.Split(strings.Split(distance_string, "Distance: ")[1], "")

	times := parseNumbers(time_split)
	distances := parseNumbers(distance_split)

	races := make([]Race, 0)

	for i := 0; i < len(times); i++ {
		races = append(races, Race{times[i], distances[i]})
	}

	single_time_string := strings.ReplaceAll(strings.Split(time_string, "Time: ")[1], " ", "")
	single_distance_string := strings.ReplaceAll(strings.Split(distance_string, "Distance: ")[1], " ", "")
	single_time, _ := strconv.Atoi(single_time_string)
	single_distance, _ := strconv.Atoi(single_distance_string)

	return races, Race{single_time, single_distance}
}

func simulateRace(race Race) []int {
	further := make([]int, 0)
	for i := 0; i <= race.duration; i++ {
		speed := i
		remaining := race.duration - i
		traveled := speed * remaining

		if traveled > race.distance {
			further = append(further, speed)
		}
	}

	return further
}
