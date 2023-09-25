package model

type Status int

const (
	Sent Status = iota
	Delivered
)

type MsgInfo struct {
	Id        int
	From      int
	To        int
	Msg       interface{}
	Status Status
}
