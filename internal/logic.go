package internal

import (
	"fmt"
	"io"
	"math"
	"slices"
	"time"
	"yadro-intership/internal/models"
	"yadro-intership/pkg/utils"

	"github.com/pkg/errors"
)

type tempTables struct {
	TimeStart time.Time
}

func Run(in io.Reader, out io.Writer) (models.TableMap, error) {
	var err error
	state, err := ParseState(in)
	if err != nil {
		return models.TableMap{}, err
	}
	var (
		Clients         = []string{}
		TablesToClients = make(map[int]string, state.Computers)
		WaitingQueue    = make([]string, 0, state.Computers)
		Tables          = models.TableMap{
			Map: make(map[int]models.Table, state.Computers),
		}
	)
	var TempTables = make(map[int]tempTables, state.Computers)
	fmt.Fprintln(out, state.StartTime.Format("15:04"))

	for err != io.EOF {
		event, err := ParseEventOnState(in, state)
		if !event.Time.IsZero() {
			if event.TableID == 0 {
				fmt.Fprintln(out, event.Time.Format("15:04"), event.ID, event.ClientName)
			} else {
				fmt.Fprintln(out, event.Time.Format("15:04"), event.ID, event.ClientName, event.TableID)
			}
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			switch err.(type) {
			case *models.NotOpenYetError:
				fmt.Fprintln(out, event.Time.Format("15:04"), int(models.OutError), err.Error())
				continue
			default:
				return models.TableMap{}, err
			}
		}

		switch event.ID {
		case int(models.InClientCame):
			if utils.Contains(Clients, event.ClientName) {
				fmt.Fprintln(out, event.Time.Format("15:04"), int(models.OutError), models.NewYouShallNotPassError().Error())
			} else {
				Clients = append(Clients, event.ClientName)
			}
		case int(models.InClientTakeAPlace):
			if !utils.Contains(Clients, event.ClientName) {
				fmt.Fprintln(out, event.Time.Format("15:04"), int(models.OutError), models.NewClientUnknownError().Error())
			}
			if v, ok := TablesToClients[event.TableID]; !ok {
				TablesToClients[event.TableID] = event.ClientName
				TempTables[event.TableID] = tempTables{
					TimeStart: event.Time,
				}
			} else if v != event.ClientName {
				fmt.Fprintln(out, event.Time.Format("15:04"), int(models.OutError), models.NewPlaceIsBusyError().Error())
			}
		case int(models.InClientWaiting):
			if len(TablesToClients) < state.Computers {
				fmt.Fprintln(out, event.Time.Format("15:04"), int(models.OutError), models.NewICanWaitNoLongerError().Error())
			} else if len(WaitingQueue) == state.Computers {
				fmt.Fprintln(out, event.Time.Format("15:04"), int(models.OutCLientLeft), event.ClientName)
			} else {
				WaitingQueue = append(WaitingQueue, event.ClientName)
			}
		case int(models.InCLientLeft):
			if !utils.Contains(Clients, event.ClientName) {
				fmt.Fprintln(out, event.Time.Format("15:04"), int(models.OutError), models.NewClientUnknownError().Error())
			}
			var table int
			for k, v := range TablesToClients {
				if v == event.ClientName {
					table = k
					Clients = utils.RemoveElem(Clients, event.ClientName)
				}
			}
			Tables.Map[table] = models.Table{
				TimeInWork: Tables.Map[table].TimeInWork + event.Time.Sub(TempTables[table].TimeStart),
				Margin:     Tables.Map[table].Margin + int(math.Ceil(event.Time.Sub(TempTables[table].TimeStart).Hours()))*state.HourlyRate,
			}
			if len(WaitingQueue) == 0 {
				delete(TempTables, table)
				continue
			}
			TablesToClients[table] = WaitingQueue[0]
			WaitingQueue = WaitingQueue[1:]
			TempTables[table] = tempTables{
				TimeStart: event.Time,
			}
			fmt.Fprintln(out, event.Time.Format("15:04"), int(models.OutClientTakeAPlace), TablesToClients[table], table)
		}
	}

	if len(Clients) > 0 {
		slices.Sort(Clients)
		for _, client := range Clients {
			fmt.Fprintln(out, state.EndTime.Format("15:04"), int(models.OutCLientLeft), client)
		}
	}
	if len(TempTables) > 0 {
		for k, v := range TempTables {
			Tables.Map[k] = models.Table{
				TimeInWork: Tables.Map[k].TimeInWork + state.EndTime.Sub(v.TimeStart),
				Margin:     Tables.Map[k].Margin + int(math.Ceil(state.EndTime.Sub(v.TimeStart).Hours()))*state.HourlyRate,
			}
		}
	}

	fmt.Fprintln(out, state.EndTime.Format("15:04"))

	return Tables, nil
}
