package models

type Event struct {
	ID        int    `json:"id"`
	NameEvent string `json:"name"`
	Shape     string `json:"shape"`
	Place     string `json:"place"`
	BeginTime string `json:"begin_time"`
	Duration  string `json:"duration"`
	//listMembers []Contact `json:"members"`
}
