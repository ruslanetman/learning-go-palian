package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 float64
	fmt.Println("Enter three numbers via space:")
	fmt.Scanf("%f %f %f", &num1, &num2, &num3)
	if num1 == num2 && num2 == num3 {
		fmt.Println("All numbers are equal.")
		return
	}
	fmt.Printf("Minimum: %.2f\n", min(num1, num2, num3))
	fmt.Printf("Average: %.2f\n", avg(num1, num2, num3))
	fmt.Printf("Maximum: %.2f\n", max(num1, num2, num3))
}
func min(a, b, c float64) float64 {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
func avg(a, b, c float64) float64 {
	return float64(a+b+c) / 3.0
}
func max(a, b, c float64) float64 {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}
