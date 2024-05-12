package internal

import (
	"fmt"
	"io"
	"strconv"
	"strings"
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

	var parsedTime, body string

	if _, err := fmt.Fscan(stream, &parsedTime, &event.ID, &body); err != nil {
		return models.Event{}, err
	}
	fmt.Println(parsedTime, event.ID, body)

	event.Time, err = time.Parse("15:04", parsedTime)
	if err != nil {
		return models.Event{}, err
	} else if event.Time.Before(state.StartTime) || event.Time.After(state.EndTime) {
		return models.Event{
			Time: event.Time,
		}, models.NewNotOpenYetError()
	}

	switch models.IncomingEvents(event.ID) {
	case models.InClientCame, models.InClientWaiting, models.InCLientLeft:
		event.ClientName = body
	case models.InClientTakeAPlace:
		splitedBody := strings.Split(body, " ")
		event.ClientName = splitedBody[0]
		event.TableID, err = strconv.Atoi(splitedBody[1])
		if err != nil {
			return models.Event{}, err
		}
	}

	return event, nil
}
