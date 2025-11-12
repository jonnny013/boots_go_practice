package main

import (
	"fmt"
	"time"
)

func main() {
	snapCh := make(chan time.Time, 10)
	saveCh := make(chan time.Time)
	logCh := make(chan string)

	go func() {
		tickerInterval := 200 * time.Millisecond
		for {
			time.Sleep(tickerInterval)
			select {
			case snapCh <- time.Now():
			default:

			}
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		saveCh <- time.Now()
		close(saveCh)
	}()

	go saveBackups(snapCh, saveCh, logCh)

	for m := range logCh {
		fmt.Println(m)
	}

}

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot(logChan)
		case <-saveAfter:
			saveSnapshot(logChan)
			return
		default:
			waitForData(logChan)
			time.Sleep(500 * time.Millisecond)
		}
	}

}

// don't touch below this line

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}


// next step: https://www.boot.dev/lessons/1f2da05c-9dda-4759-a237-74fab3dac89a