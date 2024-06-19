package main

func addEmailsToQueue(emails []string) chan string {
	ch_emails := make(chan string, len(emails))
	for _, email := range emails {
		ch_emails <- email
	}
	return ch_emails
}
