package main

import "fmt"

func main() {
	data := []int{5, 8, 2, 9, 4, 23, 10, 3}
	bubbleSort(data, greater)
	fmt.Println(data)

	data2 := []int{44, 1, 23, 45, 3, 7, 6}
	bubbleSort2(data2, greater)
	fmt.Println(data2)

	dataInsertion := []int{1, 55, 2, 4, 9, 8, 3, 6}
	insertionSort(dataInsertion, greater)
	fmt.Println("insertionSort", dataInsertion)
}

func less(val1, val2 int) bool {
	return val1 < val2
}

func greater(val1, val2 int) bool {
	return val1 > val2
}

// Bubble Sort

/*
BubbleSort is the slowest algorithm for sorting, it is used or can be used when the dataset is small

1. The outer loop  represents the number of swaps that are done for the comparison of data.
2. The inner loop is used for the comparison of data. In the first iteration, the largest value will be moved to the end of the array, in the second iteration
the second-largest value will be moved before the largest value and so on.
3.The greater() function is used for comparison within the loop and in the main function -> in the bubbleSort() method, we call the param
comp that is a callback function that takes in two ints and returns a bool, which represents the greater function and less function.

	-> O(n^2) worst case
*/
func bubbleSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	for i := 0; i < (size - 1); i++ {
		// one index less than i
		for j := 0; j < size-i-1; j++ {
			// Compare and swap
			if comp(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

// Improved bubbleSort

/*
When there is no more swapping in one pass of the outer loop, the array is already sorted. At this point, we should stop sorting
This sorting improvment in bubblesort is particuraly useful when we know that, except for a few elements, the rest of the array is already sorted.

We use the swapped variable to track whether we need to swap elements, it will be adding 1 to it if the element arr[j] is greater than arr[j+1]

	when it isnt it will be left at 0

This allows us to check if some elements are sorted already, turning our algorithm for bubbleSort()

-> O(n) because if it is nearly sorted it would not enter the second for loop and only go into the first loop iteration
*/
func bubbleSort2(arr []int, comp func(int, int) bool) {
	size := len(arr)
	swapped := 1

	for i := 0; i < (size-1) && swapped == 1; i++ {
		swapped = 0
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swapped = 1
			}
		}
	}
}

// Insertion Sort
/*
 Insertion sort works similar to how we organise a deck of cards. We keep a sorted subarray.

1. The outer loop is used to choose the value to be inserted into the sorted array on the left.
2. The chosen value we want to insert is saved in a temp variable.
3. The inner loop is doing the comparison using the greater() function. The values are shifted to the right until we find the proper position
of the temp value, for which this iteration is performed
4. Finally, temp's value is placed in its proper position, In each iteration of the outer loop, the length of the sorted aray increases by one. When we ext the outer loop
the array is sorted

Complexity: -> O(n^2)
*/

func insertionSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	var temp, i, j int
	for i = 1; i < size; i++ {
		temp = arr[i]
		for j = i; j > 0 && comp(arr[j-1], temp); j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}
