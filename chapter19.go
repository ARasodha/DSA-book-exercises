/*
Question 3: Create a function to reverse an array that takes O(1) extra space
*/

package main

import "fmt"

func reverse(numbers []int) {
	for i:=0; i<len(numbers) / 2; i++ {
		numbers[i], numbers[len(numbers) - 1 - i] = numbers[len(numbers) - 1 - i], numbers[i]
	}
}

func main() {
	numbers := []int{1,2,3,4}
	reverse(numbers)
	fmt.Println(numbers)
}