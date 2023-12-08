package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Robnarok/AOC-2023-GO/helper"
)

func main() {
	s, err := helper.ReadFile("01")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	sum := 0
	for _, v := range s {
		s2 := returnNumbers(v)
		sum += getTwoNumbers(s2)
	}
	fmt.Printf("Question 1: %v\n", sum)
	sum = 0
	for _, v := range s {
		v = mapWrittenToNumber(v)
		s2 := returnNumbers(v)
		sum += getTwoNumbers(s2)
	}
	fmt.Printf("Question 2: %v\n", sum)

}

func returnNumbers(input string) string {
	numbers := "0123456789"
	retValue := ""
	for _, s := range input {
		if strings.ContainsRune(numbers, s) {
			retValue += string(s)
		}
	}
	return retValue
}

// Just mapping the Words to the Number will cause problems. eightwo should be replaced to 82, not 8wo ..
// Because all none number Characters are going to get removed, this works...Even if its coursed and i hate it
func mapWrittenToNumber(input string) string {
	m := make(map[string]string)
	m["one"] = "one1one"
	m["two"] = "two2two"
	m["three"] = "three3three"
	m["four"] = "four4four"
	m["five"] = "five5five"
	m["six"] = "six6six"
	m["seven"] = "seven7seven"
	m["eight"] = "eight8eight"
	m["nine"] = "nine9nine"

	for k, v := range m {
		input = strings.ReplaceAll(input, k, v)
	}
	return input
}

func getTwoNumbers(input string) int {
	if len(input) == 0 {
		return 0
	}

	retValue := ""
	retValue += string(input[0])
	retValue += string(input[len(input)-1])
	i, err := strconv.Atoi(retValue)
	if err != nil {
		return 0
	}
	return i
}
