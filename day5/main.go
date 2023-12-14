package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Soil struct {
	start        int
	end          int
	source_start int
	source_end   int
	length       int
	offset       int
}

type Action struct {
	name    string
	numbers map[int][]Soil
}

type SeedRange struct {
	start  int
	end    int
	length int
}

type Field struct {
	seeds       *[]int
	seed_ranges *[]SeedRange
	actions     []Action
}

func main() {
	file, err := os.Open("part1.txt")

	if err != nil {
		log.Fatal(err)
		return
	}

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	field := Field{}
	field.actions = make([]Action, 0)
	actionLines := make([]string, 0)

	for i := 0; i < len(lines); i++ {
		if i == 0 {
			seeds, seed_ranges := parseSeeds(&lines[i])
			field.seeds = seeds
			field.seed_ranges = seed_ranges
			continue
		}

		if lines[i] == "" {
			if field.actions != nil && len(actionLines) > 0 {
				action := paseAction(&actionLines)
				field.actions = append(field.actions, *action)
				actionLines = make([]string, 0)
			}
			continue
		} else {
			actionLines = append(actionLines, lines[i])
		}
	}

	if field.actions != nil && len(actionLines) > 0 {
		action := paseAction(&actionLines)
		field.actions = append(field.actions, *action)
		actionLines = make([]string, 0)
	}

	for i, v := range *field.seeds {
		value := v
		for _, action := range field.actions {
			for _, soil := range action.numbers {
				for _, s := range soil {
					if value >= s.source_start && value <= s.source_end {
						value = value + s.offset
						goto skip
					}
				}
			}

		skip:
		}

		(*field.seeds)[i] = value
	}

	lowest := (*field.seeds)[0]
	for _, v := range *field.seeds {
		if v < lowest {
			lowest = v
		}
	}
	fmt.Println("Lowest: ", lowest)

	lowest = (*field.seed_ranges)[0].start
	for _, v := range *field.seed_ranges {
		num := lowestSeed(v, field.actions, 0)
		if num < lowest {
			lowest = num
		}
	}
	fmt.Println("Lowest: ", lowest)
}

func lowestSeed(seed SeedRange, actions []Action, start int) int {
	cr := seed
	lowest := cr.start
	for i, action := range actions {
		if i < start {
			continue
		}
		for _, soil := range action.numbers {
			for _, s := range soil {
				if cr.start >= s.source_start && cr.end <= s.source_end {
					cr.start = cr.start + s.offset
					cr.end = cr.end + s.offset
					goto skip2
				}

				if cr.start >= s.source_start && cr.start <= s.source_end && cr.end > s.source_end {
					outside_start := s.source_end + 1
					outside_end := cr.end
					outside_length := outside_end - outside_start + 1

					deepCheck := lowestSeed(SeedRange{start: outside_start, end: outside_end, length: outside_length}, actions, i)
					if deepCheck < lowest {
						lowest = deepCheck
					}

					cr.start = cr.start + s.offset
					cr.end = s.source_end + s.offset
					cr.length = cr.end - cr.start + 1
					goto skip2
				}

				if cr.start < s.source_start && cr.end >= s.source_start && cr.end <= s.source_end {
					outside_start := cr.start
					outside_end := s.source_start - 1
					outside_length := outside_end - outside_start + 1

					deepCheck := lowestSeed(SeedRange{start: outside_start, end: outside_end, length: outside_length}, actions, i)
					if deepCheck < lowest {
						lowest = deepCheck
					}

					cr.start = s.source_start + s.offset
					cr.end = cr.end + s.offset
					cr.length = cr.end - cr.start + 1
					goto skip2
				}

				if cr.start < s.source_start && cr.end > s.source_end {
					outside_start := cr.start
					outside_end := s.source_start - 1
					outside_length := outside_end - outside_start + 1

					deepCheck := lowestSeed(SeedRange{start: outside_start, end: outside_end, length: outside_length}, actions, i)
					if deepCheck < lowest {
						lowest = deepCheck
					}

					outside_start = s.source_end + 1
					outside_end = cr.end
					outside_length = outside_end - outside_start + 1

					deepCheck = lowestSeed(SeedRange{start: outside_start, end: outside_end, length: outside_length}, actions, i)
					if deepCheck < lowest {
						lowest = deepCheck
					}

					cr.start = s.source_start + s.offset
					cr.end = s.source_end + s.offset
					cr.length = cr.end - cr.start + 1
					goto skip2
				}
			}
		}

	skip2:
	}
	if cr.start < lowest {
		lowest = cr.start
	}

	return lowest
}

func paseAction(actionLines *[]string) *Action {
	action := Action{}
	action.numbers = make(map[int][]Soil)
	for i := 0; i < len(*actionLines); i++ {
		if i == 0 {
			action.name = (*actionLines)[i]
			continue
		}

		numbers := strings.Split((*actionLines)[i], " ")
		action.numbers[i] = make([]Soil, 0)
		soil := Soil{}
		for i, v := range numbers {
			num, err := strconv.Atoi(v)
			if err != nil {
				continue
			}

			if i == 0 {
				soil.start = num
			} else if i == 1 {
				soil.source_start = num
			} else if i == 2 {
				soil.length = num
			}
		}

		soil.end = soil.start + soil.length - 1
		soil.source_end = soil.source_start + soil.length - 1
		soil.offset = soil.start - soil.source_start

		action.numbers[i] = append(action.numbers[i], soil)
		soil = Soil{}
	}

	return &action
}

func parseSeeds(input *string) (*[]int, *[]SeedRange) {
	split := strings.Split(*input, ": ")
	line := split[1]
	numberSplit := strings.Split(line, " ")

	seed := make([]int, 0)
	seed_ranges := make([]SeedRange, 0)
	for _, v := range numberSplit {
		num, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		seed = append(seed, num)

		if len(seed)%2 == 0 && len(seed) > 0 {
			start := seed[len(seed)-2]
			length := seed[len(seed)-1]
			end := start + length - 1

			seed_ranges = append(seed_ranges, SeedRange{start: start, end: end, length: length})
		}
	}

	return &seed, &seed_ranges
}
