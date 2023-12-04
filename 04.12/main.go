package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input")

	cards := strings.Split(string(data), "\n")
	var result int
	for _, card := range cards {
		input := strings.Split(card, ":")
		i := strings.Split(input[1], "|")
		win := i[0]
		your := i[1]
		cardValue := 0
		for _, v := range strings.Split(your, " ") {
			if inList(strings.Split(win, " "), v) {
				if cardValue == 0 {
					cardValue = 1
				} else {
					cardValue *= 2
				}
			}
		}
		result += cardValue
	}
	fmt.Println()
	fmt.Println(result)
}

func inList(win []string, v string) (isInList bool) {
	isInList = false
	for _, w := range win {
		_, err := strconv.Atoi(w)
		if err != nil {
			continue
		}
		if v == w {
			isInList = true
			break
		}
	}

	return
}
