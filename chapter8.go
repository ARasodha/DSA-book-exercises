// Chapter 8 Exercises:
/*
Question 1:
write a function that returns the intersection of two arrays. The intersection is a third
array that contains all values contained within the first two arrays. For example, the intersection
of [1,2,3,4,5] and [0,2,4,6,8] is [2,4].
Your function should have a complexity of O(N). 
Do not use a built in way of doing this. 
*/

package main

import ( 
	"fmt"
	"regexp"
)

func intersection(arr1, arr2 []int) []int{
	arr1Vals := make(map[int]bool)
	result := make([]int, 0)

	for _, val := range arr1 {
		arr1Vals[val] = true
	}

	for _, val := range arr2 {
		if arr1Vals[val] {
			result = append(result, val)
		}
	}

	return result
}

/*
Question 2:
write a function that accepts an array of strings and returns the first duplicate value it finds.
For example, if the array is ["a", "b", "c", "d", "c", "e", "f"], the function should return "c"
since its value is duplicated within the array. 
You can assume there is one pair of duplicates in the array
Make sure the efficiency is O(N)
*/ 

func duplicates(arr []string) string {
	tracker := make(map[string]bool) 
	for _, val := range arr {
		if tracker[val] {
			return val
		}

		tracker[val] = true
	}

	return "No duplicates"
}

/*
Question 3:
write a function that accepts a string that contains all the letters of the alphabet except one
and returns the missing letter. For example, the string "the quick brown box jumps over a lazy dog"
contains all the letters except for "f".
The time complexity of the function should be O(N)
*/

func missingLetter(str string) string {
	tracker := make(map[string]bool)
	for i := range str {
		r, _ := regexp.Compile("[a-z]")
		if r.MatchString(string(str[i])) {
			tracker[string(str[i])] = true
		}
	}

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for i := range alphabet {
		if !tracker[string(alphabet[i])] {
			return string(alphabet[i])
		}
	}

	return "No missing letters"
}

/*
Question 4:
write a function that returns the first non-duplicated character in a string. For example,
the string "minimum" has two characters that only exist once, "n" and "u", so your function
should return "n" since it occurs first. The function should have an efficiency of O(N)
*/

func noDups(str string) string {
	tracker := make(map[string]int) 
	

	for i := range str {
		if tracker[string(str[i])] > 0 {
				tracker[string(str[i])] += 1
		} else {
			tracker[string(str[i])] = 1
		}
	}

	for key := range tracker {
		if tracker[key] == 1 {
			return key
		}
	}

	return "no single vals"
}

func main() {
	fmt.Println(intersection([]int{1,2,3,4,5}, []int{0,2,4,6,8})) // [2,4]
	fmt.Println(duplicates([]string{"a", "b", "c", "d", "c", "e", "f"})) // "c"
	fmt.Println(missingLetter("the quick brown box jumps over a lazy dog")) // "f"
	fmt.Println(noDups("minimum")) // "n"
}