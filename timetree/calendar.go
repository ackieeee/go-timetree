package timetree

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
)

type Calendar struct {
	httpMethod httpMethod
}

func (c *Calendar) List(includes []string) (*CalendarsData, error) {
	spath := "/calendars"
	query := QueryParam{
		&url.Values{},
	}
	if len(includes) != 0 {
		value := strings.Join(includes, ",")
		query.Set("include", value)
	}
	res, err := c.httpMethod.Get(spath, nil, &query)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data := CalendarsData{}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Calendar) Get(id string, includes []string) (*CaldendarData, error) {
	spath := "/calendars/" + id
	query := QueryParam{
		&url.Values{},
	}
	if len(includes) != 0 {
		value := strings.Join(includes, ",")
		query.Set("include", value)
	}
	res, err := c.httpMethod.Get(spath, nil, &query)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data := CaldendarData{}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Calendar) Labels(id string) (*CalendarsData, error) {
	spath := "/calendars/" + id + "/labels"
	query := QueryParam{
		&url.Values{},
	}
	res, err := c.httpMethod.Get(spath, nil, &query)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data := CalendarsData{}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Calendar) UpcomingEvents(id string, timezone string, days string, includes []string) (*CalendarsData, error) {
	spath := "calendars/" + id + "/upcoming_events"
	if id == "" {
		return nil, errors.New("id must not be empty")
	}
	if timezone == "" {
		return nil, errors.New("timezone must not be empty")
	}
	if days == "" {
		return nil, errors.New("days must not be empty")
	}

	query := QueryParam{
		&url.Values{},
	}
	if len(includes) != 0 {
		value := strings.Join(includes, ",")
		query.Set("include", value)
	}

	res, err := c.httpMethod.Get(spath, nil, &query)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data := CalendarsData{}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
