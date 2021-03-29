package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type (
	sendRequestOptions struct {
		URL                       string
		Method                    string
		ContentType               string
		UserName                  string
		Password                  string
		UserFieldName             string
		PasswordFieldName         string
		UserAgent                 string
		Timeout                   time.Duration
		ForceLoginFieldsIntoQuery bool
		IgnoreHTTPErrorCodes      bool
	}
)

func sendRequest(options sendRequestOptions) (resp *http.Response, err error) {
	client := &http.Client{
		Timeout: options.Timeout * time.Second,
	}

	newURL, err := mergeLoginFieldsIntoHTTPQuery(options)

	if err != nil {
		return
	}

	var bodyContent io.Reader
	formData := url.Values{}
	if options.Method != "GET" && !options.ForceLoginFieldsIntoQuery {
		formData.Set(options.UserFieldName, options.UserName)
		formData.Set(options.PasswordFieldName, options.Password)
		bodyContent = strings.NewReader(formData.Encode())
	}

	req, err := http.NewRequest(options.Method, newURL, bodyContent)

	if err != nil {
		return nil, fmt.Errorf("NewRequest: %s", err)
	}

	req.Header.Set("User-Agent", options.UserAgent)

	if options.ContentType != "" {
		req.Header.Add("Content-Type", options.ContentType)
	}

	if bodyContent != nil {
		req.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))
	}

	resp, err = client.Do(req)

	if err != nil {
		return
	}

	if !options.IgnoreHTTPErrorCodes {
		if resp.StatusCode != http.StatusOK {
			return resp, fmt.Errorf("http: status code given was %d", resp.StatusCode)
		}
	}

	return
}

func mergeLoginFieldsIntoHTTPQuery(options sendRequestOptions) (outURL string, err error) {

	// Nothing to do. So just return the url supplied in the options since
	if options.Method != "GET" && !options.ForceLoginFieldsIntoQuery {
		outURL = options.URL
		return
	}

	u, err := url.Parse(options.URL)

	if err != nil {
		return
	}

	q := u.Query()
	q.Set(options.UserFieldName, options.UserName)
	q.Set(options.PasswordFieldName, options.Password)
	u.RawQuery = q.Encode()

	outURL = u.String()

	return
}
