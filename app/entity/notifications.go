package entity

type Notification interface {
	Isnotification()
}

type Basenotification struct {
	Createdat int64
}

func (Basenotification) Isnotification() {}

type Messagenotification struct {
	Basenotification
	Noty string
}

/*
type Unreadworkrequest struct {
	Basenotification
	Workid int    `json:"workid"`
	Title  string `json:"title"`
}

type Unreadmessagenotification struct {
	Basenotification
	Count int `json:"count"`
}
*/
