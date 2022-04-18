package main

import "fmt"
/*
Question 1:
Use recursion to write a function that accepts an array of strings wand returns the total number
of characters across all strings. For example, of the input array is ["ab", "c", "def", "ghij"],
the output should be 10 since there are 10 characters total
*/

func countChars(slice []string) int {
	if len(slice) == 1 {
		return len(slice[0])
	}

	return len(slice[0]) + countChars(slice[1:])
}

/*
Question 2:
Use recursion to write a function that accepts an array of numbers and returns a new array
containing just the event numbers
*/

func even(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	if numbers[0] % 2 == 0 {
		return append([]int{numbers[0]}, even(numbers[1:])...)
	} else {
		return even(numbers[1:])
	}
}

/*
Question 3:
There is a numerical sequence known as "Triangular Numbers". The pattern begins as 1,3,6,10,15,21
and continues onward with the Nth number in the pattern, which is N plus previous number.
For example, the 7th number in the sequence is 28, since it's 7 (which is N) plus 21 (the previous 
number in the sequence). Write a function that accepts a number for N and returns the correct
number from the series. That is, if the function is passed 7, the function should return 28
*/

func triangle(n int) int {
	if n == 1 {
		return 1
	}

	return n + triangle(n - 1)
}

/*
Question 4:
Use recursion to write a function that accepts a string and returns the first index that contains
the character "x"/ For example, "abcdefghijklmnopqrstuvwx" has an "x" at index 23. To keep things
simple assume the input string has at least one "x".
*/

func findX(str string) int {
	if str[0] == 'x' {
		return 0
	}

	return findX(str[1:]) + 1
}


/*
// Question 5:
This problem is known as the "Unique Paths" problem: Let's say you have a grid of rows and columns.
Write a function that accepts a number of columns and calculates the number of possible "shortest"
paths from the upper-left most square to the lower right-most square.
*/

func uniquePaths(rows, columns int) int {
	if rows == 1 || columns == 1 {
		return 1
	}

	return uniquePaths(rows - 1, columns) + uniquePaths(rows, columns - 1)
}

func main() {
	fmt.Println(countChars([]string{"ab", "c", "def", "ghij"}))	// 10
	fmt.Println(even([]int{1,2,3,4,5,6,7,8,9,10})) // [2,4,6,8,10]
	fmt.Println(evenNums([]int{1,2,3,4,5,6,7,8,9,10}, 0)) // [2,4,6,8,10]
	fmt.Println(triangle(7)) // 28
	fmt.Println(findX("abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(uniquePaths(3,7)) // 
}