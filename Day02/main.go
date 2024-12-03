package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToInt(stringSlice []string) []int {
	intSlice := make([]int, len(stringSlice))

	for i, s := range stringSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		intSlice[i] = num
	}
	return intSlice
}

func isLevelSafe(l []int) bool {
	// if level is decreasing
	if l[0] > l[1] {
		for i := 0; i < len(l)-1; i++ {
			diff := l[i] - l[i+1]
			if diff == 1 || diff == 2 || diff == 3 {
				continue
			} else {
				return false
			}
		}
		return true
		// if level is increasing
	} else if l[0] < l[1] {
		for i := 0; i < len(l)-1; i++ {
			diff := l[i+1] - l[i]
			if diff == 1 || diff == 2 || diff == 3 {
				continue
			} else {
				return false
			}
		}
		return true
	}
	return false
}

func removeInt(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func isLevelReallySafe(l []int) bool {
	copy1 := make([]int, len(l))
	copy(copy1, l)
	copy2 := make([]int, len(l))
	copy(copy2, l)
	copy3 := make([]int, len(l))
	copy(copy3, l)
	// if level is decreasing
	if l[0] > l[1] {
		for i := 0; i < len(l)-1; i++ {
			diff := l[i] - l[i+1]
			if diff == 1 || diff == 2 || diff == 3 {
				continue
			} else {
				copy1 = removeInt(copy1, i)
				if i != 0 {
					copy2 = removeInt(copy2, i-1)
				}
				if i < len(l)+1 {
					copy3 = removeInt(copy3, i+1)
				}

				try1 := isLevelSafe(copy1)
				try2 := isLevelSafe(copy2)
				try3 := isLevelSafe(copy3)
				if try1 || try2 || try3 {
					return true
				}
				return false
			}
		}
		return true

		// if level is increasing
	} else if l[0] < l[1] {
		for i := 0; i < len(l)-1; i++ {
			diff := l[i+1] - l[i]
			if diff == 1 || diff == 2 || diff == 3 {
				continue
			} else {

				copy1 = removeInt(copy1, i)
				if i != 0 {
					copy2 = removeInt(copy2, i-1)
				}
				if i < len(l)+1 {
					copy3 = removeInt(copy3, i+1)
				}

				try1 := isLevelSafe(copy1)
				try2 := isLevelSafe(copy2)
				try3 := isLevelSafe(copy3)
				if try1 || try2 || try3 {
					return true
				}
				return false
			}
		}
		return true
	}
	// if level is the same at the beginning
	newL := removeInt(l, 0)
	return isLevelSafe(newL)
}
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input.txt", err)
		return
	}
	defer file.Close()

	safe := 0
	reallySafe := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		level := strings.Split(scanner.Text(), " ")
		intSlice := stringToInt(level)
		if isLevelSafe(intSlice) {
			safe++
		}
		if isLevelReallySafe(intSlice) {
			reallySafe++
		}

	}
	fmt.Println("Total number safe are: ", safe)
	fmt.Println("Total number really safe are: ", reallySafe)
}
