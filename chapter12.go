package main

import ( 
	"fmt"
	"strconv"
)

/*
Question 1:
The following function accepts an array of numbers and returns the sum, as long as a particular
number doesn't bring the sum above 100. If adding a particular number will make the sum higher
than 100, the number is ignored. However, this function makes unnecessary recursive calls. Fix
the code to eliminate the unnecessary recursion.package book
*/

// func addUntil100(arr []int) int {
// 	if len(arr) == 0 {
// 		return 0
// 	}

// 	if arr[0] + addUntil100(arr[1:]) > 100 {
// 		return addUntil100(arr[1:])
// 	} else {
// 		return arr[0] + addUntil100(arr[1:])
// 	}
// }

func addUntil100(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	till100 := addUntil100(arr[1:])

	if arr[0] + till100 > 100 {
		return till100
	} else {
		return arr[0] + till100
	}
}


/*
Question 2:
The following function uses recursion to calculate the Nth number from a mathematical sequence 
known as the "Golomb Sequence". It's terribly inefficient, though! Use memoization to optimize it!
(You don't actually have to understand how Golomb Sequence works to do this exercise.)
*/

// func golomb(n int) int {
// 	if n == 1 {
// 		return 1
// 	}

// 	return 1 + golomb(n - golomb(golomb(n - 1)))
// }

func golomb(n int, memo map[int]int) int {
	if n == 1 {
		return 1
	}

	_, contained := memo[n]

	if !contained {
		memo[n] = 1 + golomb(n - golomb(golomb(n - 1, memo), memo), memo)
	}

	return memo[n]
}

/*
Question 3:
Here is a solution to "Unique Paths" problem from the previous chapter exercises. Yse memoization
to improve its efficiency.
*/

// func uniquePaths(rows, columns int) int {
// 	if rows == 1 || columns == 1 {
// 		return 1
// 	}

// 	return uniquePaths(rows - 1, columns) + uniquePaths(rows, columns - 1)
// }

func uniquePaths(rows, columns int, memo map[string]int) int {
	if rows == 1 || columns == 1 {
		return 1
	}

	path := strconv.Itoa(rows) + strconv.Itoa(columns)
	_, contained := memo[path]

	if !contained {
		memo[path] = uniquePaths(rows - 1, columns, memo) + uniquePaths(rows, columns - 1, memo)
	}

	return memo[path]
 
}

func main() {
	// fmt.Println(addUntil100([]int{20, 30, 25, 26, 25})) // 96
	// fmt.Println(golomb(5, map[int]int{})) // 3
	fmt.Println(uniquePaths(3, 7, map[string]int{})) // 24
}