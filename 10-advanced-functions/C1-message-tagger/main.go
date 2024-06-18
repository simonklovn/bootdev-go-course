package main

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	taggedMessages := []sms{}
	for _, message := range messages {
		message.tags = tagger(message)
		taggedMessages = append(taggedMessages, message)
	}
	return taggedMessages
}

func tagger(msg sms) []string {
	tags := []string{}
	if strings.Contains(msg.content, "urgent") || strings.Contains(msg.content, "Urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(msg.content, "sale") || strings.Contains(msg.content, "Sale") {
		tags = append(tags, "Promo")
	}
	return tags
} 
