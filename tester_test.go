package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestExecute (t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r	*http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("the data"))
	}))
	defer ts.Close()

	subject := Tester{
		Url:ts.URL,
		Data:"the data",
	}

	result, err := subject.Execute()

	if err != nil {
		t.Error(err)
	}

	if result.Response != "the data" {
		t.Errorf("\nExpected: 'the data'\n Got: %v", result.Response)
	}

	if result.Status != 200 {
		t.Errorf("\nExpected: 200\n Got: %v", result.Status)
	}
}
