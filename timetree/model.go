package timetree

type CalendarData struct {
	Data []Item `json:"data"`
}

type Item struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Order       int    `json:"order"`
	ImageURL    string `json:"image_url"`
	CreatedAt   string `json:"created_at"`
}
