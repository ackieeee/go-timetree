package timetree

type CalendarsData struct {
	Data []Item `json:"data"`
}

type CaldendarData struct {
	Data Item `json:"data"`
}

type Item struct {
	ID            string        `json:"id"`
	Type          string        `json:"type"`
	Attributes    Attributes    `json:"attributes"`
	Relationships Relationships `json:"relationships"`
}

type Attributes struct {
	Title         string `json:"title"`
	AllDay        string `json:"all_day"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Color         string `json:"color"`
	Order         int    `json:"order"`
	ImageURL      string `json:"image_url"`
	StartAt       string `json:"start_at"`
	StartTimezone string `json:"start_timezone"`
	EndAt         string `json:"end_at"`
	EndTimezone   string `json:"end_timezone"`
	Location      string `json:"location"`
	URL           string `json:"url"`
	UpdatedAt     string `json:"updated_at"`
	CreatedAt     string `json:"created_at"`
	Category      string `json:"category"`
	Recurrence    string `json:"recurrence"`
	RecurringUUID string `json:"recurring_uuid"`
}

type Relationships struct {
	Labels LabelsData `json:"labels"`
}

type LabelsData struct {
	Data []Label `json:"data"`
}

type Label struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
