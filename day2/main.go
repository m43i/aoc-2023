package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
    Color string
    Amount int
}

type Round struct {
    Number int
    Cubes []Cube
    Red int
    Blue int
    Green int
}

func (r *Round) IsValid() bool {
    max_red := 12
    max_blue := 14
    max_green := 13

    if r.Red > max_red {
        return false
    }

    if r.Blue > max_blue {
        return false
    }

    if r.Green > max_green {
        return false
    }

    return true
}

type Game struct {
    Id int
    Rounds []Round
    Red int
    Blue int
    Green int
}

func (g *Game) IsValid() bool {
    valid := true

    for _, round := range g.Rounds {
        if !round.IsValid() {
            valid = false
            break
        }
    }

    return valid
}

func main() {
    file, err := os.Open("part1.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    valid_games := 0
    result := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        value := scanner.Text()
        game := parseGame(&value)

        if game.IsValid() {
            valid_games += game.Id
        }

        product := game.Red * game.Blue * game.Green
        result += product
    }

    fmt.Println("Valid games: ", valid_games)
    fmt.Println("Result: ", result)
}

func parseCube(value *string) Cube {
    split := strings.Split(*value, " ")
    amount, _ := strconv.Atoi(split[0])
    color := split[1]

    return Cube{
        Color: color,
        Amount: amount,
    }
}

func parseRound(value *string, num int) Round {
    split := strings.Split(*value, ", ")

    var cubes []Cube
    for _, cube := range split {
        cubes = append(cubes, parseCube(&cube))
    }

    red := 0
    blue := 0
    green := 0

    for _, cube := range cubes {
        if cube.Color == "red" {
            red += cube.Amount
        } else if cube.Color == "blue" {
            blue += cube.Amount
        } else if cube.Color == "green" {
            green += cube.Amount
        }
    }

    return Round{
        Number: num,
        Cubes: cubes,
        Red: red,
        Blue: blue,
        Green: green,
    }
}

func parseGame(value *string) Game {
    split := strings.Split(*value, ": ")
    game_split := split[0]
    game_id, _ := strconv.Atoi(strings.Split(game_split, " ")[1]);

    var rounds []Round
    rounds_split := strings.Split(split[1], "; ")
    for i, round := range rounds_split {
        rounds = append(rounds, parseRound(&round, i+1))
    }

    red := -1
    blue := -1
    green := -1

    for _, round := range rounds {
        if red == -1 {
            red = round.Red
        } else if red < round.Red {
            red = round.Red
        }

        if blue == -1 {
            blue = round.Blue
        } else if blue < round.Blue {
            blue = round.Blue
        }

        if green == -1 {
            green = round.Green
        } else if green < round.Green {
            green = round.Green
        }
    }

    return Game{
        Id: game_id,
        Rounds: rounds,
        Red: red,
        Blue: blue,
        Green: green,
    }
}

