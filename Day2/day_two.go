package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getNum(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return -1
}

type CubeSet struct {
	red, blue, green int
}

func (r CubeSet) isPossible(c CubeSet) bool {
	if r.red <= c.red && r.blue <= c.blue && r.green <= c.green {
		return true
	}
	return false
}

type CubeSetBuilder struct {
	CubeSet CubeSet
}

func NewCubeSetBuilder() *CubeSetBuilder {
	return &CubeSetBuilder{
		CubeSet: CubeSet{},
	}
}
func (gb *CubeSetBuilder) Build() CubeSet {
	return gb.CubeSet
}

type Game struct {
	id   int
	sets []CubeSet
}

func (r Game) validGame(c CubeSet) (bool, int) {
	for _, set := range r.sets {
		if !set.isPossible(c) {
			return false, 0
		}
	}
	return true, r.id
}

func (r Game) powerMinGame() int {
	minR := 0
	minG := 0
	minB := 0
	for _, set := range r.sets {
		if set.red > minR {
			minR = set.red
		}
		if set.green > minG {
			minG = set.green
		}
		if set.blue > minB {
			minB = set.blue
		}
	}
	return minB * minG * minR
}

type GameBuilder struct {
	Game Game
}

func NewGameBuilder() *GameBuilder {
	return &GameBuilder{
		Game: Game{},
	}
}

func (gb *GameBuilder) setGame(input string) {
	gameSplit := strings.SplitAfter(input, ":")
	game := gameSplit[0]
	gb.Game.id = getNum(game[5 : len(game)-1])
	sets := gameSplit[1]
	splitSets := strings.Split(sets, ";")
	for _, set := range splitSets {
		cs := NewCubeSetBuilder()
		colours := strings.Split(set, ",")
		for _, combo := range colours {
			combo = strings.TrimLeft(combo, " ")
			comboArr := strings.Split(combo, " ")
			count := getNum(comboArr[0])
			colour := comboArr[1]
			if colour == "red" {
				cs.CubeSet.red = count
			} else if colour == "green" {
				cs.CubeSet.green = count
			} else {
				cs.CubeSet.blue = count
			}
			gb.Game.sets = append(gb.Game.sets, cs.Build())
		}
	}
	//fmt.Println(gb.Game.sets)

}

func (gb *GameBuilder) Build() Game {
	return gb.Game
}

func sumValidGames(games []Game, valid CubeSet) int {
	count := 0
	for _, game := range games {
		isValid, gameId := game.validGame(valid)
		if isValid {
			count += gameId
		}
	}
	return count
}

func sumMinPowers(games []Game) int {
	count := 0
	for _, game := range games {
		power := game.powerMinGame()
		//fmt.Println(power)
		count += power

	}
	return count
}

func main() {
	file, err := os.Open("Day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	MaxSet := CubeSet{red: 12, green: 13, blue: 14}
	var allGames []Game

	for scanner.Scan() {
		game := NewGameBuilder()
		game.setGame(scanner.Text())
		allGames = append(allGames, game.Build())
	}
	fmt.Println(sumValidGames(allGames, MaxSet))
	fmt.Println(sumMinPowers(allGames))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
