package main

func countReports(numSentCh chan int) int {
	var total int 
	for {
		numSent, ok := <-numSentCh
		if !ok {
			break
		}
		total += numSent
	}
	return total
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
