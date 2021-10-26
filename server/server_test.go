package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//testing getword count function

func TestGetWordCount(t *testing.T) {
	///*
	req, err := http.NewRequest("GET", "localshost:9000/count", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CountWords)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := " 14"
	data, _ := ioutil.ReadAll(rr.Body)

	var a []byte
	a = append(a, 49, 52, 10)
	fmt.Println(a, data)
	if string(a) != string(data) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
