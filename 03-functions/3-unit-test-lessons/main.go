package main

func getMonthlyPrice(tier string) int {
	const penniesInADollar int = 100;
	switch tier {
	case "basic":
		return 100 * penniesInADollar
	case "premium":
		return 150 * penniesInADollar
	case "enterprise":
		return 500 * penniesInADollar
	default:
		return 0
	}
}