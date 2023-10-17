package main

import (
	"fmt"
	"http-load-tester/Infrastructure"
	"time"
)

func main() {
	Infrastructure.Run(func(cts Infrastructure.CancellationTokenSource) {
		requestInfo, err := loadRunDetails()
		if err != nil {
			fmt.Println(err)
			return
		}
		startTicker(requestInfo, 5, cts)
	})
}

func startTicker(info *HttpRequestInfo, totalSeconds uint16, cts Infrastructure.CancellationTokenSource) {
	ticker := time.NewTicker(time.Second * 5)

	httpRequest, responseChannel, err := SetupRequest(info)
	if err != nil {
		fmt.Println(err)
	}

	secondsRemaining := totalSeconds
	go func() {
		defer close(responseChannel)

		for {
			select {
			case <-ticker.C:
				secondsRemaining -= 1
				SendRequests(info, responseChannel, httpRequest)
				if secondsRemaining == 0 {
					ticker.Stop()
					fmt.Printf("%d seconds elapsed. Collecting results.\n", totalSeconds)
					close(cts)
					return
				}
			case <-cts:
				ticker.Stop()
				fmt.Println("Ticker stopped.")
				return
			}
		}
	}()
}
