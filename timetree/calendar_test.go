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

func TestCalendar_Get(t *testing.T) {
	t.Helper()
	tj, err := os.Open("../testdata/calendar_get.json")
	if err != nil {
		t.Fatal(err)
	}

	testHttpCli := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       tj,
			Header:     make(http.Header),
		}
	})

	token := "test"
	cli := NewClient(testHttpCli, token)
	data, err := cli.Calendar.Get("1", []string{})
	if err != nil {
		t.Fatal(err)
	}
	if data == nil {
		t.Fatalf("unexpected result. data is nil\n")
	}
}

func TestCalendar_Labels(t *testing.T) {
	t.Helper()
	tj, err := os.Open("../testdata/calendar_labels.json")
	if err != nil {
		t.Fatal(err)
	}

	testHttpCli := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       tj,
			Header:     make(http.Header),
		}
	})
	token := "test"
	cli := NewClient(testHttpCli, token)
	data, err := cli.Calendar.Labels("1")
	if err != nil {
		t.Fatal(err)
	}
	if data == nil {
		t.Fatalf("unexpected result. data is nil\n")
	}
	if data.Data[0].Type != "label" {
		t.Fatalf("unexpected result. type name is not label\n")
	}
}
func TestCalendar_UpcomingEvents(t *testing.T) {
	t.Helper()
	tj, err := os.Open("../testdata/calendar_upcoming_events.json")
	if err != nil {
		t.Fatal(err)
	}

	testHttpCli := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       tj,
			Header:     make(http.Header),
		}
	})
	token := "test"
	cli := NewClient(testHttpCli, token)
	data, err := cli.Calendar.UpcomingEvents("1", "Asia/Tokyo", "3", []string{"creator", "label", "attendees"})
	if err != nil {
		t.Fatal(err)
	}
	if data == nil {
		t.Fatalf("unexpected result. data is nil\n")
	}
	if data.Data[0].Type != "event" {
		t.Fatalf("unexpected result. type name is not event\n")
	}
}
