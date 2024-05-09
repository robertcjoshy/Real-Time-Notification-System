package entity

import "time"

type Notification interface {
	Isnotification()
}

type Basenotification struct {
	Createdat time.Time `json:"createdat"`
}

func (Basenotification) Isnotification() {}

type Unreadworkrequest struct {
	Basenotification
	Workid int    `json:"workid"`
	Title  string `json:"title"`
}

type Unreadmessagenotification struct {
	Basenotification
	Count int `json:"count"`
}
