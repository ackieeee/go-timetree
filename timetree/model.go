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
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Order       int    `json:"order"`
	ImageURL    string `json:"image_url"`
	CreatedAt   string `json:"created_at"`
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
