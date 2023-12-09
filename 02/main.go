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

func main() {
	_, err := helper.ReadFile("023")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func parseGames(input string) []string {
	s := strings.Split(input, ":")
	if len(s) != 2 {
		return nil
	}
	return strings.Split(s[1], ";")
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
