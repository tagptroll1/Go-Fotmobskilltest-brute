package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type void struct{}

func getInput(text string) string {
	var input string
	fmt.Print(text)
	fmt.Scanln(&input)
	return input
}

func getRangeNumbers(rangeStr string) []int {
	var rangeNumbersStr []string = strings.Split(rangeStr, "-")
	var start, end int
	start, _ = strconv.Atoi(rangeNumbersStr[0])
	end, _ = strconv.Atoi(rangeNumbersStr[1])
	var numbers []int
	for i := start; i <= end; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func getSliceOfNumbersFromStringRange(rangeStr string) map[int]void {
	setOfNumber := make(map[int]void)
	var rangeNumbersStr []string = strings.Split(rangeStr, ",")

	for _, rangeNumberStr := range rangeNumbersStr {
		if strings.Contains(rangeNumberStr, "-") {
			var rangeNumbers []int = getRangeNumbers(rangeNumberStr)
			for _, rangeNumber := range rangeNumbers {
				setOfNumber[rangeNumber] = void{}
			}
		} else {
			var number int
			number, _ = strconv.Atoi(rangeNumberStr)
			setOfNumber[number] = void{}
		}
	}
	return setOfNumber
}

// Get all numbers in includes, and exclude all numbers in excludes
func getNumbers(includes, excludes string) []int {
	var numbers []int
	includeRanges := getSliceOfNumbersFromStringRange(includes)
	excludeRanges := getSliceOfNumbersFromStringRange(excludes)

	for excludeRange := range excludeRanges {
		delete(includeRanges, excludeRange)
	}

	for includeRange := range includeRanges {
		numbers = append(numbers, includeRange)
	}

	sort.Ints(numbers)
	return numbers
}

func formatNumbers(numbers []int) string {
	var numbersStr []string
	var lastNumber int

	for i, number := range numbers {
		if i == 0 {
			numbersStr = append(numbersStr, fmt.Sprintf("%d", number))
			lastNumber = number
		} else if i == len(numbers)-1 {
			numbersStr = append(numbersStr, fmt.Sprintf("-%d", number))
		} else if lastNumber == number-1 {
			lastNumber = number
		} else {
			numbersStr = append(numbersStr, fmt.Sprintf("-%d, %d", lastNumber, number))
			lastNumber = number
		}
	}

	return strings.Join(numbersStr, "")
}

func main() {
	includes := getInput("Includes: ")
	excludes := getInput("Excludes: ")

	numbers := getNumbers(includes, excludes)
	fmt.Printf("Numbers: %s\n", formatNumbers(numbers))
}
