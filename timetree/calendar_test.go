package timetree

import (
	"net/http"
	"os"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestCalendar_List(t *testing.T) {
	t.Helper()
	tj, err := os.Open("../testdata/calendar_list.json")
	if err != nil {
		t.Fatal(err)
	}
	defer tj.Close()

	testHttpCli := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       tj,
			Header:     make(http.Header),
		}
	})

	token := "test"
	cli := NewClient(testHttpCli, token)
	data, err := cli.Calendar.List([]string{})
	if err != nil {
		t.Fatal(err)
	}
	if data == nil {
		t.Fatalf("unexpected result. data is nil\n")
	}
}
