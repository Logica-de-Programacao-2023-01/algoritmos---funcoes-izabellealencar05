package main

// Crie uma função que receba um slice de inteiros e retorne o segundo maior valor.
func secondLargest(numbers []int) int {
	if len(numbers) < 2 {
		return 0
	}

	largest := numbers[0]
	secondLargest := numbers[1]

	if secondLargest > largest {
		largest, secondLargest = secondLargest, largest
	}

	for _, n := range numbers[2:] {
		if n > largest {
			secondLargest = largest
			largest = n
		} else if n > secondLargest && n != largest {
			secondLargest = n
		}
	}

	return secondLargest
}
