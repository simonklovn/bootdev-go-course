package main

func concurrentFib(n int) []int {
	values := []int{}
	ch_fib := make(chan int)

	go func() {
		fibonacci(n, ch_fib)
	}()

	for item := range ch_fib {
		values = append(values, item)
	}
	return values
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
