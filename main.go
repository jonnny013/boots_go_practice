package main

import (
	"fmt"
	"time"
)

type email struct {
	body string
	date time.Time
}




func main() {
		runCases :=
		[3]email{
			{
				body: "Words are pale shadows of forgotten names. As names have power, words have power.",
				date: time.Date(2019, 2, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "It's like everyone tells a story about themselves inside their own head.",
				date: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "Bones mend. Regret stays with you forever.",
				date: time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
			},
		}
	result := checkEmailAge(runCases)
	fmt.Println(result)
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}
