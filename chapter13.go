package main

import (
	"fmt"
	"sort"
)

/*
Question 1: 
Given an array of positive numbers, write a function that returns the greatest product of any three
numbers. The approach of using three nested loops would clock in at O(N^3), which is very slow.
Use sorting to implement the function in a way that computes a O(N log N) speed. (There are even
faster implementations, but we're focusing on using sorting as technique to make code faster.)
*/

func greatestProduct(arr []int) int {
	sort.Ints(arr)
	greatestThree := arr[len(arr) - 3:]
	sum := 0

	for _, num := range greatestThree {
		sum += num
	}
	
	return sum
}

/*
Question 2:
The following function finds the "missing number" from an array of integers. That is, the array
is expected to have all integers from 0 and up to the array's length, but one is missing. As
examples, the array, [5,2,4,1,0] is missing the number 3, and the array [9,3,2,5,6,7,1,0,4] is 
missing the number 8.
Here's an implementation of O(N^2) (the `includes` method alone is already O(N), since the computer
needs to search the entire array to find n):
*/

// Provided initial implementation not copied from textbook: uses for loop on unsorted array
// and includes method in JavaScript to check val of index in loop if it wasn't in array on each
// round of iteration

func missingNum(numbers []int) int {
	sort.Ints(numbers)
	for i, num := range numbers {
		if i != num {
			return i
		}
	}

	return 0
}

/*
Question 3: 
Write three different implementations of a function that fines the greatest number in an array.
Write one that is O(N^2), one that is O(N log N) and one that is O(N)
*/

// O(N^2)
func greatest1(numbers []int) int {
	for i:=0; i < len(numbers); i++ {
		temp := numbers[i]
		current := numbers[i]
		for j:= i + 1; j < len(numbers); j++ {
			if numbers[i] < numbers[j] {
				current = numbers[j]
			}
		}

		if temp == current {
			return temp
		}
	}

	return 0
}

// O(N log N)
func greatest2(numbers []int) int {
	sort.Ints(numbers)
	return numbers[len(numbers) - 1]
}

// O(N)
func greatest3(numbers []int) int {
	greatest := numbers[0]
	for _, val := range numbers {
		if val > greatest {
			greatest = val
		}
	}

	return greatest
}

func main() {
	fmt.Println(greatestProduct([]int{2,3,6,1,5})) // 14
	fmt.Println(missingNum([]int{5,2,4,1,0}))      // 3
	fmt.Println(missingNum([]int{9,3,2,5,6,7,1,0,4})) // 8
	fmt.Println(greatest1([]int{9,3,2,5,6,7,1,0,4})) // 9
	fmt.Println(greatest2([]int{9,3,2,5,6,7,1,0,4})) // 9
	fmt.Println(greatest3([]int{9,3,2,5,6,7,1,0,4})) // 9
}