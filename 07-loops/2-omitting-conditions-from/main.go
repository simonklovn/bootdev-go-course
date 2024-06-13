package main

func maxMessages(thresh int) int {
	const messageCost = 100
	const feeIncrease = 1

	additionalFee 	:= 0
	totalCost 		:= 0

	for numMessages := 0; ; numMessages++ {
		nextCost := messageCost+additionalFee
		if totalCost + nextCost > thresh {
			return numMessages
		}
		totalCost += nextCost
		additionalFee += feeIncrease
	}
}
