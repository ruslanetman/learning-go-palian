package main

import (
	"fmt"
	"strconv"
	"strings"
)

// This program calculates the minimum, average, and maximum of a list of numbers

func main() {
	fmt.Println("Enter numbers separated by spaces:")
	var input string
	fmt.Scanln(&input)

	parts := strings.Fields(input)
	if len(parts) == 0 {
		fmt.Println("No numbers entered.")
		return
	}

	numbers := make([]float64, 0, len(parts))
	for _, p := range parts {
		num, err := strconv.ParseFloat(p, 64)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", p)
			return
		}
		numbers = append(numbers, num)
	}

	min, max, sum := numbers[0], numbers[0], 0.0
	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
		sum += n
	}
	avg := sum / float64(len(numbers))

	fmt.Printf("Minimum: %.2f\n", min)
	fmt.Printf("Average: %.2f\n", avg)
	fmt.Printf("Maximum: %.2f\n", max)
}
