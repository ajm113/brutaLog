package main

import (
	"net/http"
	"testing"
	"time"
)

func TestSendRequestPOST(t *testing.T) {
	resp, err := sendRequest(sendRequestOptions{
		URL:               "https://reqbin.com/echo/post/form",
		Method:            "POST",
		ContentType:       "application/x-www-form-urlencoded",
		UserFieldName:     "user",
		PasswordFieldName: "pass",
		UserName:          "andrew",
		Password:          "mypassword",
		UserAgent:         "my-test-go-client",
		Timeout:           2 * time.Second,
	})

	if err != nil {
		t.Error("failed sending request for following:", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Error("expected status code is OK, instead got:", resp.StatusCode)
	}
}

func TestSendRequestGET(t *testing.T) {
	resp, err := sendRequest(sendRequestOptions{
		URL:               "https://reqbin.com/echo/get/json",
		Method:            "GET",
		UserFieldName:     "user",
		PasswordFieldName: "pass",
		UserName:          "andrew",
		Password:          "mypassword",
		UserAgent:         "my-test-go-client",
		Timeout:           2 * time.Second,
	})

	if err != nil {
		t.Error("failed sending request for following:", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Error("expected status code is OK, instead got:", resp.StatusCode)
	}
}
