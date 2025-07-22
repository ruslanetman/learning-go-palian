package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter numbers separated by space (press Ctrl+D on Linux/Mac or Ctrl+Z on Windows when done):")
	var nums []float64
	for {
		var x float64
		n, err := fmt.Scan(&x)
		if n == 0 || err != nil {
			break
		}
		nums = append(nums, x)
	}
	if len(nums) < 2 {
		fmt.Println("Enter at least 2 numbers.")
		return
	}

	fmt.Printf("Minimum: %.2f\n", min(nums...))
	fmt.Printf("Average: %.2f\n", avg(nums...))
	fmt.Printf("Maximum: %.2f\n", max(nums...))
}

func min(nums ...float64) float64 {
	m := nums[0]
	for _, v := range nums[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func avg(nums ...float64) float64 {
	sum := 0.0
	for _, v := range nums {
		sum += v
	}
	return sum / float64(len(nums))
}

func max(nums ...float64) float64 {
	m := nums[0]
	for _, v := range nums[1:] {
		if v > m {
			m = v
		}
	}
	return m
}
