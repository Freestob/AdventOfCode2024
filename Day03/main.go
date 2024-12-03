package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func stringToInt(s []string) []int {
	intSlice := make([]int, len(s))

	for i, s := range s {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		intSlice[i] = num
	}
	return intSlice
}
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input.txt", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	total := 0
	for scanner.Scan() {
		level := scanner.Text()
		matches := re.FindAllStringSubmatch(level, -1)
		for _, match := range matches {
			digits := stringToInt(match[1:])
			total = total + (digits[0] * digits[1])
		}
	}
	fmt.Println(total)
}
