package models

// IncomingEvents
type IncomingEvents int

const (
	InClientCame IncomingEvents = iota + 1
	InClientTakeAPlace
	InClientWaiting
	InCLientLeft
)

func (e IncomingEvents) String() string {
	return [...]string{
		"ClientCame",
		"ClientTakeAPlace",
		"ClientWaiting",
		"ClientLeft",
	}[e-1]
}

// OutcomingEvents
type OutcomingEvents int

const (
	OutCLientLeft OutcomingEvents = iota + 11
	OutClientTakeAPlace
	OutError
)

func (e OutcomingEvents) String() string {
	return [...]string{
		"ClientLeft",
		"ClientTakeAPlace",
		"Error",
	}[e-11]
}
