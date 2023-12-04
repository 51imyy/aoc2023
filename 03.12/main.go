package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input")
	data := string(file)
	lineLength := len(strings.Split(data, "\n")[0])
	var savedIndexes []int
	var result []int
	list := []int{
		-1,              // l
		1,               // r
		-lineLength - 1, // t
		-lineLength - 2, // tl
		-lineLength,     // tr
		lineLength,      // bl
		lineLength + 1,  // b
		lineLength + 2,  // b
	}

	rp := regexp.MustCompile("[^\\d.\n]")
	matches := rp.FindAllStringIndex(data, -1)
	for _, match := range matches {
		symbolIndex := match[0]
		// p2: per match save times found int
		// if times == 2 multiply the numbers
		// add than this to result list
		// sum result list
		for _, index := range list {
			a := string(data[symbolIndex+index])
			if !isDigit(a) {
				continue
			}
			lIndex := index
			IntStr := a
			for symbolIndex+lIndex-1 >= 0 && isDigit(string(data[symbolIndex+lIndex-1])) {
				IntStr = fmt.Sprintf("%s%s", string(data[symbolIndex+lIndex-1]), IntStr)
				lIndex--
			}

			rIndex := index
			for symbolIndex+rIndex+1 < len(data) && isDigit(string(data[symbolIndex+rIndex+1])) {
				IntStr = fmt.Sprintf("%s%s", IntStr, string(data[symbolIndex+rIndex+1]))
				rIndex++
			}

			if !contains(savedIndexes, symbolIndex+lIndex) {
				i, _ := strconv.Atoi(IntStr)
				result = append(result, i)
				savedIndexes = append(savedIndexes, symbolIndex+lIndex)
			}
		}
	}

	fmt.Println(sum(result))
}

func isDigit(str string) (isDigit bool) {
	_, err := strconv.Atoi(str)
	isDigit = err == nil
	return
}

func contains(indexList []int, index int) (contains bool) {
	contains = false
	for _, i := range indexList {
		if i == index {
			contains = true
			break
		}
	}

	return
}

func sum(list []int) (result int) {
	for _, i := range list {
		result += i
	}

	return
}
