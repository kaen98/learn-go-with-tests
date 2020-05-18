package main

func Sum(numbers [5]int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}