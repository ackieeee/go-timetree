package timetree

import (
	"encoding/json"
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
