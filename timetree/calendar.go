package timetree

import "encoding/json"

type Calendar struct {
	httpMethod httpMethod
}

func (c *Calendar) List() (*CalendarData, error) {
	spath := "/calendars"
	res, err := c.httpMethod.Get(spath, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data := CalendarData{}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
