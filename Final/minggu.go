// buat generate slice dan selection sort

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	slice := generateSlice(20)
	p(slice, "\n")

	// BubbleSort
	tB_start := time.Now()
	p("Bubble Sort : \n", tB_start)
	selectionSort(slice)
	tB_end := time.Now()
	tB_elapsed := tB_end.Sub(tB_start)
	p(slice, tB_elapsed, "\n\n")
	// -----------------

	// SelectionSort
	tS_Start := time.Now()
	p("Selection Sort : \n", tS_Start)
	bubbleSort(slice)
	tS_end := time.Now()
	tElapsed := tS_end.Sub(tS_Start)
	p(slice, tElapsed, "\n\n")
	// -----------------

	// InsertionSort
	tI_start := time.Now()
	p("Insertion : \n", tI_start)
	insertionSort(slice)
	tI_end := time.Now()
	tI_elapsed := tI_end.Sub(tI_start)
	p(slice, tI_elapsed, "\n\n")
	// -----------------

}

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&slice[i])
		if slice[i]%2 == 1 {

		} else {
			i = i - 1
			fmt.Println("Oops Only accept Odd numbers")
		}
	}
	return slice
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
		j := 1
		if j > 0 {
			if items[j-1] > items[i] {
				items[j-1], items[i] = items[i], items[j-1]
			}
		}
		j = j - 1
	}
}
