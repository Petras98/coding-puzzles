package main

import "fmt"

/*
remove duplicates when multiple are present - not sorted
*/


func removeDuplicates(nums []int) []int{
	var values map[int]bool = make(map[int]bool)
	var finalArr []int = make([]int, 0)

	//O(n)
	for _, val := range nums{
		_, exists := values[val]
		if !exists{
			values[val] = true
			finalArr = append(finalArr, val)
		}
	}

	return finalArr
}

func main(){
	
	nums := []int{2, 8, 5, 8, 2, 44, 5 }

	uniqueArr := removeDuplicates(nums)

	fmt.Println(uniqueArr)
}