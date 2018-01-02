package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	numbers := generateArray(20)
	p(numbers, "\n\n")
	p("Sorting ... \n\n")

	// SelectionSort
	s_Start := time.Now()
	selectionSort(numbers)
	p("\n\n SELECTION SORT==> ")
	p("Start ", s_Start, "\n", numbers)
	s_End := time.Now()
	p("End ", s_End)
	s_Elapsed := s_End.Sub(s_Start)
	p("Total time :", s_Elapsed)
	// End of SelectionSort

	// BubbleSort
	b_Start := time.Now()
	bubbleSort(numbers)
	p("\n\n BUBBLE SORT==> ")
	p("Start ", b_Start, "\n", numbers)
	b_End := time.Now()
	p("End ", b_End)
	b_Elapsed := b_End.Sub(b_Start)
	p("Total time :", b_Elapsed)
	// End of BubbleSort

	// InsertionSort
	i_Start := time.Now()
	insertionSort(numbers)
	p("\n\n INSERTION SORT==> ")
	p("Start ", i_Start, "\n", numbers)
	i_End := time.Now()
	p("End ", i_End)
	i_Elapsed := i_End.Sub(i_Start)
	p("Total time :", i_Elapsed)
	// End of InsertionSort
}

func generateArray(size int) []int {

	array := make([]int, size, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&array[i])
		if array[i]%2 == 1 {

		} else {
			i = i - 1
			fmt.Println("Opps, only accept only odd number ")
		}
	}
	return array

}

func selectionSort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}

func bubbleSort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}

func insertionSort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
