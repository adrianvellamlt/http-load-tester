package main

import (
	"bufio"
	"fmt"
	infra "http-load-tester/Infrastructure"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type HttpRequestInfo struct {
	requestsPerSecond uint16
	uri               *url.URL
	authScheme        *string
	authToken         *string
	bodyContents      *string
	httpClient        http.Client
}

func loadRunDetails() (*HttpRequestInfo, error) {
	requestInfo := HttpRequestInfo{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Specify the desired requests per second:")
	rpsStr, _, _ := reader.ReadLine()
	rps, err := strconv.ParseInt(string(rpsStr), 10, 16)
	if err != nil {
		return &requestInfo, err
	} else {
		requestInfo.requestsPerSecond = uint16(rps)
	}

	fmt.Println("Specify a URI to POST Requests to:")
	urlStr, _, _ := reader.ReadLine()
	uri, err := url.ParseRequestURI(string(urlStr))
	if err != nil {
		return &requestInfo, err
	} else {
		requestInfo.uri = uri
	}

	fmt.Println("Specify the Auth header (with scheme) if applicable:")
	authHeader, _, _ := reader.ReadLine()
	if len(authHeader) > 0 {
		authHeaderSplit := strings.Split(string(authHeader), " ")
		if len(authHeaderSplit) != 2 ||
			(!strings.EqualFold(authHeaderSplit[0], "basic") &&
				!strings.EqualFold(authHeaderSplit[0], "bearer")) {
			e := fmt.Errorf("%s is not a valid Auth Header. Format should be: [Scheme (Basic/Bearer)] [Token]", string(authHeader))
			return &requestInfo, e
		} else {
			requestInfo.authScheme = infra.StringPtr(authHeaderSplit[0])
			requestInfo.authToken = infra.StringPtr(authHeaderSplit[1])
		}
	} else {
		requestInfo.authScheme = nil
		requestInfo.authToken = nil
	}

	fmt.Println("Specify the Request Body Contents:")
	contents, _, _ := reader.ReadLine()
	requestInfo.bodyContents = infra.StringArrPtr(contents)

	return &requestInfo, nil
}
