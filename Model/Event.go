package Model

type Event struct {
	Timestamp string `json:"timestamp"`
	VisitorId int    `json:"visitorId"`
	Event     string `json:"event"`
	ItemId    int    `json:"ItemId"`
}
