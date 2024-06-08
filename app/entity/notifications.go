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
