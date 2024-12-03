package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.
func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		fmt.Println("Error opening input.txt", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var s1 []float64
	var s2 []float64
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "   ")
		num1 := stringToInt(nums[0])
		num2 := stringToInt(nums[1])
		s1 = append(s1, float64(num1))
		s2 = append(s2, float64(num2))
	}
	sort.Float64s(s1)
	sort.Float64s(s2)

	total := float64(0)
	for i := 0; i < len(s1); i++ {
		diff := math.Abs(s1[i] - s2[i])
		total += diff
	}
	fmt.Println("Total = ", total)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
