package helper

import (
	"log"
	"os"
	"strings"
)

func ReadFile(filename string) ([]string, error) {
	content, err := os.ReadFile("input/" + filename + ".txt")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	s := strings.Split(string(content), "\n")
	return s, nil
}
