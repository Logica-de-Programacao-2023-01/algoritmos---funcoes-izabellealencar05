package main

func findIndex(numbers []int, value int) int {
	for i, num := range numbers {
		if num == value {
			return i
		}
	}
	return -1
}
