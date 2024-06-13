package main

func sum(nums ...int) int {
	sumOfInputs := 0
	for i := 0; i <len(nums); i++ {
		sumOfInputs += nums[i]
	}
	return sumOfInputs
}