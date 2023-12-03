package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var dict = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func main() {
	file, err := os.Open("part1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var nums []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num := parseLine(&line, false)
		nums = append(nums, num)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var sum int
	for _, num := range nums {
		sum += num
	}

	println(sum)
}

func parseLine(s *string, scheck bool) int {
	var nums map[int]string = make(map[int]string)
	for i, word := range dict {
		if scheck {
			index := strings.Index(*s, word)
			lastIndex := strings.LastIndex(*s, word)
			if index != -1 {
				nums[index] = strconv.Itoa(i + 1)
			}
			if index != lastIndex && lastIndex != -1 {
				nums[lastIndex] = strconv.Itoa(i + 1)
			}
		}

		idx := strings.Index(*s, strconv.Itoa(i+1))
		lastIdx := strings.LastIndex(*s, strconv.Itoa(i+1))
		if idx != -1 {
			nums[idx] = strconv.Itoa(i + 1)
		}
		if idx != lastIdx && lastIdx != -1 {
			nums[lastIdx] = strconv.Itoa(i + 1)
		}
	}

	smallest := -1
	highest := -1
	for key := range nums {
		if smallest == -1 || key < smallest {
			smallest = key
		}
		if highest == -1 || key > highest {
			highest = key
		}
	}

	result := nums[smallest] + "" + nums[highest]
	num, _ := strconv.Atoi(result)
	return num
}
