package main

import "fmt"

func missingNumber(nums []int) int {
    n := len(nums)
    v := 0;
    var i int;
    for i = 0; i < n; i++ {
        v = v ^ i ^ nums[i] ;
    }
    v ^= n;
    return v;
    
}

func checkSolution(nums []int, expected int){
	if out := missingNumber(nums); out != expected{
		fmt.Println("Failed")
		fmt.Println("\tInput: ", nums)
		fmt.Println("\tOutput: ", out, "Expected: ", expected)
	}else{
		fmt.Println("Passed")
		fmt.Println("\tInput: ", nums)
		fmt.Println("\tOutput: ", out)
	}
}

func main(){
	nums := []int{3,0,1}
	checkSolution(nums, 2)
	
	nums = []int{0,1}
	checkSolution(nums, 2)

	nums = []int{9,6,4,2,3,5,7,0,1}
	checkSolution(nums, 8)
}