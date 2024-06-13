package main

import "fmt"

func (e email) cost() int {
	costPerChar := 5
	if e.isSubscribed {
		costPerChar = 2
	}
	return costPerChar * len(e.body)
}

func (e email) format() string {
	subscriptionStatus := "Subscribed"
	if !e.isSubscribed{
		subscriptionStatus = "Not Subscribed"
	}
	return fmt.Sprintf("'%s' | %s", e.body, subscriptionStatus)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
	expense 	 int
	formatter 	 string
}