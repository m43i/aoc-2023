package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Symbol struct {
	Line     int
	Position int
	Value    string
}

type Number struct {
	Id    int
	Line  int
	Start int
	End   int
	Value int
}

func (n *Number) IsNeighbour(line int, pos int) bool {
	if (line >= n.Line-1 && line <= n.Line+1) && (pos >= n.Start-1 && pos <= n.End) {
		return true
	}

	return false
}

func main() {
	file, err := os.Open("part1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	symbols := make([]Symbol, 0)
	numbers := make([]Number, 0)
	idx := 0
	for scanner.Scan() {
		value := scanner.Text()
		parseLine(&value, &idx, &numbers, &symbols)
		idx++
	}

	gears := make(map[string][]Number, 0)
	result := 0
	for _, number := range numbers {
		for _, symbol := range symbols {
			if number.IsNeighbour(symbol.Line, symbol.Position) {
				result += number.Value
				if symbol.Value == "*" {
					key := strconv.Itoa(symbol.Line) + "" + strconv.Itoa(symbol.Position)
					gears[key] = append(gears[key], number)
				}
			}
		}
	}

	gearResult := 0
	for _, gear := range gears {
		if len(gear) == 2 {
			val := gear[0].Value * gear[1].Value
			gearResult += val
		}
	}

	fmt.Println(result)
	fmt.Println(gearResult)
}

func parseLine(line *string, lineIdx *int, numbers *[]Number, symbols *[]Symbol) {
	tokens := strings.Split(*line, "")
	start := -1
	end := -1
	value := ""
	for i, token := range tokens {
		_, err := strconv.Atoi(token)

		if err == nil {
			if start == -1 {
				start = i
			}
			value += token
			continue
		} else {
			if start != -1 {
				end = i
				num, _ := strconv.Atoi(value)
				number := Number{
					Line:  *lineIdx,
					Start: start,
					End:   end,
					Value: num,
				}
				*numbers = append(*numbers, number)

				start = -1
				end = -1
				value = ""
			}
			if token != "." {
				*symbols = append(*symbols, Symbol{
					Line:     *lineIdx,
					Position: i,
					Value:    token,
				})
			}
		}
	}

	if start != -1 {
		end = len(tokens)
		num, _ := strconv.Atoi(value)
		number := Number{
			Line:  *lineIdx,
			Start: start,
			End:   end,
			Value: num,
		}
		*numbers = append(*numbers, number)
	}
}
