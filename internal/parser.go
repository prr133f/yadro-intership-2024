package internal

import (
	"fmt"
	"io"
	"time"
	"yadro-intership/internal/models"
)

func ParseState(stream io.Reader) (models.State, error) {
	var state models.State
	var err error

	var timeStart, timeClose string

	fmt.Fscan(stream, &state.Computers)
	fmt.Fscan(stream, &timeStart, &timeClose)
	fmt.Fscan(stream, &state.HourlyRate)

	state.StartTime, err = time.Parse("15:04", timeStart)
	if err != nil {
		return models.State{}, err
	}

	state.EndTime, err = time.Parse("15:04", timeClose)
	if err != nil {
		return models.State{}, err
	}

	return state, nil
}

func ParseEventOnState(stream io.Reader, state models.State) (models.Event, error) {
	var event models.Event
	var err error

	var parsedTime string

	if _, err := fmt.Fscan(stream, &parsedTime, &event.ID, &event.ClientName); err != nil {
		return models.Event{}, err
	}
	if event.ID == 2 {
		if _, err := fmt.Fscan(stream, &event.TableID); err != nil {
			return models.Event{}, err
		}
	}

	event.Time, err = time.Parse("15:04", parsedTime)
	if err != nil {
		return models.Event{}, err
	} else if event.Time.Before(state.StartTime) || event.Time.After(state.EndTime) {
		return event, models.NewNotOpenYetError()
	}

	return event, nil
}
