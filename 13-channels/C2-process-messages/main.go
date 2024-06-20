package main

import "sync"

func processMessages(messages []string) []string {
	var wg sync.WaitGroup
	ch_processedMsg := make(chan string, len(messages))
	processedMessages := []string{}

	for _, msg := range messages {
		wg.Add(1)
		go func (msg string) {
			defer wg.Done()
			processedMsg := msg + "-processed"
			ch_processedMsg <- processedMsg
		}(msg)
	}

	go func() {
		wg.Wait()
		close(ch_processedMsg)
	}()

	for processedMsg := range ch_processedMsg {
		processedMessages = append(processedMessages, processedMsg)
	}
	return processedMessages
}
