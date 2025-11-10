package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func getMessageText(a *Analytics, m Message) {
	if m.Success {
		a.MessagesSucceeded = a.MessagesSucceeded + 1
		a.MessagesTotal = a.MessagesTotal + 1
	} else {
		a.MessagesFailed = a.MessagesFailed + 1
				a.MessagesTotal = a.MessagesTotal + 1
	}
}
