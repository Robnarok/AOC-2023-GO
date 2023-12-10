package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Robnarok/AOC-2023-GO/helper"
)

type Cubes struct {
	red   int
	blue  int
	green int
}

func (foo *Cubes) isValid(input Cubes) bool {
	if foo.blue < input.blue {
		return false
	}
	if foo.red < input.red {
		return false
	}
	if foo.green < input.green {
		return false
	}
	return true
}

func (foo *Cubes) getPower() int {
	return foo.red * foo.blue * foo.green
}

func main() {
	input, err := helper.ReadFile("02")
	if err != nil {
		fmt.Println(err)
		return
	}
	//Only 12 red cubes, 13 green cubes, and 14 blue cubes
	validCubes := Cubes{
		red:   12,
		green: 13,
		blue:  14,
	}
	sum := 0

	for i, v := range input {
		sets := parseGames(v)

		usedCubes := []Cubes{}

		for _, sets := range sets {
			c := parseSet(sets)
			usedCubes = append(usedCubes, c)
		}
		c := maxCubes(usedCubes)

		if validCubes.isValid(c) {
			fmt.Printf("%v: %v\n", i+1, c)
			sum += i + 1
		}
	}

	fmt.Printf("Riddle One: %v\n", sum)

	sum = 0

	for _, v := range input {
		sets := parseGames(v)

		usedCubes := []Cubes{}

		for _, sets := range sets {
			c := parseSet(sets)
			usedCubes = append(usedCubes, c)
		}
		c := maxCubes(usedCubes)
		sum += c.getPower()
	}
	fmt.Printf("Riddle Two: %v\n", sum)
}

func parseGames(input string) []string {
	s := strings.Split(input, ":")
	if len(s) != 2 {
		return nil
	}
	return strings.Split(s[1], ";")
}

func maxCubes(input []Cubes) Cubes {
	max := Cubes{
		blue:  0,
		red:   0,
		green: 0,
	}

	for _, v := range input {
		if v.blue >= max.blue {
			max.blue = v.blue
		}
		if v.red >= max.red {
			max.red = v.red
		}
		if v.green >= max.green {
			max.green = v.green
		}
	}
	return max
}

func parseSet(input string) Cubes {
	s := strings.Split(input, ",")

	var cube Cubes
	cube.red = 0
	cube.blue = 0
	cube.green = 0

	for _, v := range s {
		v = strings.ReplaceAll(v, " ", "")
		if strings.Contains(v, "red") {
			v = strings.ReplaceAll(v, "red", "")
			i, err := strconv.Atoi(v)
			if err != nil {
				return Cubes{0, 0, 0}
			}
			cube.red = i
		}
		if strings.Contains(v, "blue") {
			v = strings.ReplaceAll(v, "blue", "")
			i, err := strconv.Atoi(v)
			if err != nil {
				return Cubes{0, 0, 0}
			}
			cube.blue = i
		}

		if strings.Contains(v, "green") {
			v = strings.ReplaceAll(v, "green", "")
			i, err := strconv.Atoi(v)
			if err != nil {
				return Cubes{0, 0, 0}
			}
			cube.green = i
		}
	}
	return cube

}
