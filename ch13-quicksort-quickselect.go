package main

import "fmt"

type SortableArray struct {
	arr []int
}

func (s SortableArray) Partition(leftPointer, rightPointer int) int {
	// We always choose the right-most element as the pivot
	// We keep the index of the pivot for later use:
	pivotIndex := rightPointer

	// We grab the pivot value itself:
	pivot := s.arr[pivotIndex]

	// We start the right pointer immediately to the left of the pivot
	rightPointer -= 1

	for true {
		// Move the left pointer to the right as long as it points to
		// the value that is less than the pivot
		for s.arr[leftPointer] < pivot {
			leftPointer += 1
		}

		// Move the right pointer to the left as long as it points
		// to a value that is greater than the pivot
		for s.arr[rightPointer] > pivot {
			rightPointer -= 1
		}

		// We've now reached the point where we've stopped moving
		// both the left and right pointers

		// We check whether the left pointer has reached or gone
		// beyond the right pointer. If it has, we break out of
		// the loop so we can swap the pivot later on in our code
		if leftPointer >= rightPointer {
			break
		} else {
			// If the left pointer is still to the left of the right
			// pointer, we swap the value of the left and right pointers
			s.arr[leftPointer], s.arr[rightPointer] = s.arr[rightPointer], s.arr[leftPointer]

			// We move the left pointer over to the right, gearing up
			// for the next round of the left and right pointer movements:
			leftPointer += 1
		}
	}

	// As the final step of the partition, we swap the value
	// of the left pointer with the pivot:
	s.arr[leftPointer], s.arr[pivotIndex] = s.arr[pivotIndex], s.arr[leftPointer]

	// We return the leftPointer for the sake of the quicksort method
	// which we will implement leater
	return leftPointer
}

func (s SortableArray) Quicksort(leftIndex, rightIndex int) int {
	// Base case: the subarray has 0 or 1 elements:
		if rightIndex - leftIndex <= 0 {
			return 0
		}

		// Partition the range of elements and grab the index of the pivot:
		pivotIndex := s.Partition(leftIndex, rightIndex)

		// Recursively call this Quicksort method on whatever is to the left of the pivot:
		s.Quicksort(leftIndex, pivotIndex - 1)

		// Recursively call this Quicksort method on whatever is to the right of the pivot:
		s.Quicksort(pivotIndex + 1, rightIndex)

		return 0
}

// *****Not complete, issues with returning the lowest value:*********
// func (s SortableArray) Quickselect(kthLowestValue, leftIndex, rightIndex int) int {
// 	// If we reach a base case - that is, that the subarray has one cell,
// 	// we know we've found the value we're looking for:
// 	if rightIndex - leftIndex <= 0 {
// 		return s.arr[leftIndex]
// 	}

// 	// Partition the array and grab the index of the pivot
// 	pivotIndex := s.Partition(leftIndex, rightIndex)

// 	// If what we're looking for is to the left of the pivot:
// 	if kthLowestValue < pivotIndex {
// 		// Recursively perform quickselect on the subarray to the left of the pivot:
// 		s.Quickselect(kthLowestValue, leftIndex, pivotIndex - 1)
// 		// If what we are looking for is to the right of the pivot:
// 	} else if kthLowestValue > pivotIndex {
// 		// Recursively perform quickselect on subarray to the right of the pivot:
// 		s.Quickselect(kthLowestValue, pivotIndex + 1, rightIndex)
// 		// If after the partition, the pivot position is in the same spot as the
// 		// kthLowestValue, we've found the value we are looking for
// 	} else if kthLowestValue == pivotIndex {
// 		return s.arr[pivotIndex]
// 	}
// }

func main() {
	arr := []int{0,5,2,1,6,3}
	// arr := []int{0, 50, 20, 10, 60, 30}
	s := SortableArray{arr}
	s.Quicksort(0, len(s.arr) - 1) // [0 1 2 3 6 5]
	fmt.Println(arr)
	// fmt.Println(s.Quickselect(2, 0, len(s.arr) - 1))
}