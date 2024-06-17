package main

func adder() func(int) int {
	var sum int
	return func(x int) int {
		sum += x
		return sum
	}
}
 