package models

// YouShallNotPassError
func NewYouShallNotPassError() error {
	return &YouShallNotPassError{
		msg: "YouShallNotPass",
	}
}

type YouShallNotPassError struct {
	msg string
}

func (err *YouShallNotPassError) Error() string {
	return err.msg
}

// NotOpenYet
func NewNotOpenYetError() error {
	return &NotOpenYetError{
		msg: "NotOpenYet",
	}
}

type NotOpenYetError struct {
	msg string
}

func (err *NotOpenYetError) Error() string {
	return err.msg
}

// PlaceIsBusy
func NewPlaceIsBusyError() error {
	return &PlaceIsBusyError{
		msg: "PlaceIsBusy",
	}
}

type PlaceIsBusyError struct {
	msg string
}

func (err *PlaceIsBusyError) Error() string {
	return err.msg
}

// ClientUnknown
func NewClientUnknownError() error {
	return &ClientUnknownError{
		msg: "ClientUnknown",
	}
}

type ClientUnknownError struct {
	msg string
}

func (err *ClientUnknownError) Error() string {
	return err.msg
}

// ICanWaitNoLonger!
func NewICanWaitNoLongerError() error {
	return &ICanWaitNoLongerError{
		msg: "ICanWaitNoLonger!",
	}
}

type ICanWaitNoLongerError struct {
	msg string
}

func (err *ICanWaitNoLongerError) Error() string {
	return err.msg
}
