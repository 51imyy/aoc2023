package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	spelledNumbers = []string{
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
)

func main() {
	data, _ := os.ReadFile("input2")
	lines := strings.Split(string(data), "\n")
	reg, _ := regexp.Compile("(\\d)")
	var result int

	for _, line := range lines {
		tmpLine := line
		for k, number := range spelledNumbers {
			if strings.Contains(line , number) {
				tmpLine = strings.ReplaceAll(tmpLine, number, number + strconv.Itoa(k+1) + number)
			}
		}
		regexResult := reg.FindAllString(tmpLine, -1)
		i, _ := strconv.Atoi(regexResult[0] + regexResult[len(regexResult)-1])
		result += i
	}
	
	fmt.Println("result", result)
}

type Name struct {
	Test string `json:"abc"`
}
