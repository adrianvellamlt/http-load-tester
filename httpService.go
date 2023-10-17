package main

import (
	"http-load-tester/Infrastructure"
	"net/http"
	"strings"
	"time"
)

type HttpResponseInfo struct {
	IsSuccess          bool
	RequestTimestamp   time.Time
	TimeTakenInSeconds float32
	StatusCode         int
}

func SetupRequest(info *HttpRequestInfo) (*http.Request, chan HttpResponseInfo, error) {
	bodyReader := strings.NewReader(Infrastructure.StringCoalesce(info.bodyContents, ""))
	request, err := http.NewRequest(http.MethodPost, info.uri.String(), bodyReader)

	return request, make(chan HttpResponseInfo), err
}

func SendRequests(info *HttpRequestInfo, responseChannel chan HttpResponseInfo, request *http.Request) {
	for i := 0; i < int(info.requestsPerSecond); i++ {
		go func() {

			requestStart := time.Now()
			response, err := http.DefaultClient.Do(request)
			timeTaken := time.Now().Sub(requestStart).Seconds()

			responseChannel <- HttpResponseInfo{
				IsSuccess:          err == nil,
				StatusCode:         response.StatusCode,
				RequestTimestamp:   requestStart,
				TimeTakenInSeconds: float32(timeTaken),
			}
		}()
	}
}
