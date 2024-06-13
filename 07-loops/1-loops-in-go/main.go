package main

func bulkSend(numMessages int) float64 {
	var (
		messageCost 	float64 = 1.0
		additionalCost 	float64
		totalCost 		float64
	)

	for i := 0; i < numMessages; i++ {
		totalCost += messageCost+additionalCost
		additionalCost += 0.01
	}
	return totalCost
}