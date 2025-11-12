package main

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))

	for i := range emails {
		ch <- emails[i]
	}

	return ch
}

// https://www.boot.dev/lessons/d2a10614-3142-4d3e-906e-5a817aa920b3
