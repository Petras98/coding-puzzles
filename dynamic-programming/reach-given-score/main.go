package main

import "fmt"
/*
Consider a game where a player can score 3 or 5 or 10 points in a move. Given a total score n, find number of distinct combinations to reach the given score.

Example:
Input
3
8
20
13
Output
1
1
4
2
Explanation
For 1st example when n = 8
{ 3, 5 } and {5, 3} are the 
two possible permutations but 
these represent the same 
cobmination. Hence output is 1.

*/

func findScore(n int) int{
	table := make([]int, n+1)
	table[0] = 1

	for i := 3; i <= n; i++{
		table[i] += table[i-3]
	}

	for i := 5; i <= n; i++{
		table[i] += table[i-5]
	}

	for i := 10; i <= n; i++{
		table[i] += table[i-10]
	}

	fmt.Println(table)
	return table[n]
}

func main(){
	fmt.Println(findScore(3))
	fmt.Println(findScore(8))
	fmt.Println(findScore(20))
	fmt.Println(findScore(13))

	var stocks map[string]float64
	price := stocks["RANDMO"]
	fmt.Printf("%f\n", price)

}