package main

func yearsUntilEvents(age int) (int, int, int) {
	var yearsUntilAdult int = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	var yearsUntilDrinking int = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	var yearsUntilCarRental int = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}
