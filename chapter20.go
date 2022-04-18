/*
Your're working on software that analyzes sports players. Following are two arrays of players of
different sports.package book

If you look carefully, you'll see that there are some players who participate in more than one kind 
of sport. Jill Huang and Wanda Vakulskas play both basketball and football.

You are to write a function that accepts two arrays of players and returns an array of the players
who play in both sports. In this case, that would be ["Jill Huang", "Wanda Vakulskas"]

While there are players who share first names and players who share last names, we can assume that
there's only one person who has a particular full name [first and last names].

We can use a nested-loop approach, comparing each player one array against each player from the
other array, but this would have a runtime of O(N * M).Your job is to optimize the function so that
it can run just O(N + M)
*/
package main

import "fmt"

// func players(basketballPlayers, footballPlayers []map[string]string) []string {
// 	names := make(map[string]int, 0)

// 	for _, player := range basketballPlayers {
// 		if names[player["firstName"] + " " + player["lastName"]] == 0 {
// 			names[player["firstName"] + " " + player["lastName"]] = 1
// 		} else {
// 			names[player["firstName"] + " " + player["lastName"]] += 1
// 		}
// 	}

// 	for _, player := range footballPlayers {
// 		if names[player["firstName"] + " " + player["lastName"]] == 0{
// 			names[player["firstName"] + " " + player["lastName"]] = 1
// 		} else {
// 			names[player["firstName"] + " " + player["lastName"]] += 1
// 		}
// 	}

// 	result := make([] string, 0)
// 	for player := range names {
// 		if names[player] == 2 {
// 			result = append(result, player)
// 		}
// 	}

// 	return result
// }

/*
Question 2:
You're writing a function that accepts an array of distinct integers from 0,1,2,3..up to N.
However, the array will be missing one integer and your function is to return hte missing one.
For example, this array has all the integers from 0 to 6, but is missing the 4: [2,3,0,6,1,5]
Therefore, the function should return 4. 
The next example, has all integers from 0 to 9, but is missing the 1: [0,2,3,9,4,7,5,0,6]
IN this case, the function should return 1.
Using nested loops would take up to O(N^2). Your job is to optimize the code to O(N) runtime.
*/

// func missingNumber(numbers []int) int {
// 	var tracker = make(map[int]bool, 0)

// 	for _, num := range numbers {
// 		tracker[num] = true
// 	}

// 	for i := range numbers {
// 		if !tracker[i] {
// 			return i
// 		}
// 	}

// 	return 0
// }

/*
Question 3:
Your're working on some more stock-prediction software. The function you're writing accepts an 
array of predicted prices for a particular stock over the course of time. 
For example, this array of seven prices: [10,7,5,8,11,2,6] predicts that a given stock will have
these prices over the next seven days. (On day 1, the stock will close $10, on Day , the stock will
close at $7; and so on).

Your function should calculate the greatest profit that could be made from a single "buy" transaction
folowed by a single "sell" transaction. 
In the previous example, the most money could be made if we bought the stock when it was work $5 
and sold it when it was worth $11. That yields a profit of $6 per share. 

Note that we could use nested loops to find the profit of every possible buy-and-sell combination. 
However, this would be O(N^2) and too slow for our hotshot trading platform. Your job is to optimize
the code so that the function clocks in at just O(N).
*/

func stockTrack(predictions []int) int {
	buyPrice := predictions[0]
	var greatestProfit int

	for _, price := range predictions {
		potentialProfit := price - buyPrice

		if price < buyPrice {
			buyPrice = price
		} else if potentialProfit > greatestProfit {
			greatestProfit = potentialProfit
		}
	}

	return greatestProfit
}

func main() {
	// var basketballPlayers = []map[string]string {
	// 	{"firstName": "Jill", "lastName": "Huang", "team": "Gators"},
	// 	{"firstName": "Janko", "lastName": "Barton", "team": "Sharks"},
	// 	{"firstName": "Wanda", "lastName": "Vakulskas", "team": "Sharks"},
	// 	{"firstName": "Jill", "lastName": "Moloney", "team": "Gators"},
	// 	{"firstName": "Luuk", "lastName": "Watkins", "team": "Gators"},
	// }
	
	// var footballPlayers = []map[string]string {
	// 	{"firstName": "Hanzla", "lastName": "Radosti", "team": "32ers"},
	// 	{"firstName": "Tina", "lastName": "Watkins", "team": "Barleycorns"},
	// 	{"firstName": "Alex", "lastName": "Patel", "team": "32ers"},
	// 	{"firstName": "Jill", "lastName": "Huang", "team": "Barleycorns"},
	// 	{"firstName": "Wanda", "lastName": "Vakulskas", "team": "Barleycorns"},
	// }

	// fmt.Println(players(basketballPlayers, footballPlayers))
	// fmt.Println(missingNumber([]int{2,3,0,6,1,5})) // 4
	// fmt.Println(missingNumber([]int{0,2,3,9,4,7,5,0,6})) // 1
	fmt.Println(stockTrack([]int{10,7,5,8,11,2,6})) // 6
}