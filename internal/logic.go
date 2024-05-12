package internal

import (
	"fmt"
	"io"
	"yadro-intership/internal/models"

	"github.com/pkg/errors"
)

func Run(in io.Reader) error {
	var err error
	state, err := ParseState(in)
	if err != nil {
		return errors.WithStack(err)
	}

	for err != io.EOF {
		event, err := ParseEventOnState(in, state)
		fmt.Println(event)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			switch err.(type) {
			case *models.NotOpenYetError:
				fmt.Println(err.Error())
			default:
				return err
			}
		}

		switch event.ID {
		case int(models.InClientCame):
			// TODO: Обработать происходящее событие
			// TODO: Проверить наличие места в клубе (столы + очередь)
		case int(models.InClientTakeAPlace):

		}
	}

	return nil
}
