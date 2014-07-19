package main

import (
	"errors"
	"github.com/awsmsrc/llog"
	"time"
)

type Event struct {
	Id        int       `json:"id"`
	AccountId int       `json:"account_id"`
	Data      string    `json:"data"`
	Key       string    `json:"key"`
	Attempts  *[]*Event `json:"events,omitempty"`
	State     string    `json:"state"`
}

func (e *Event) TableName() string {
	return "events"
}

func (e *Event) Attributes() ([]string, []interface{}) {
	return []string{
			"account_id",
			"data",
		}, []interface{}{
			e.AccountId,
			e.Data,
		}
}

func (e *Event) Registrations() []Registration {
	var result []Registration
	rows, err := db.Query(
		"SELECT id, url FROM registrations WHERE account_id=? AND key=?",
		e.AccountId,
		e.Key,
	)
	if err != nil {
		llog.FATAL(err)
	}

	defer rows.Close()
	for rows.Next() {
		var reg = Registration{
			AccountId: e.AccountId,
			Key:       e.Key,
		}
		if err := rows.Scan(&reg.Id, &reg.Url); err != nil {
			llog.FATAL(err)
		}
		result = append(result, reg)
	}
	if err := rows.Err(); err != nil {
		llog.FATAL(err)
	}
	return result
}

func GetEvents(account_id int) []*Event {
	var result []*Event
	rows, err := db.Query(
		"SELECT id, account_id, data, state FROM events WHERE account_id=? LIMIT 100",
		account_id,
	)
	if err != nil {
		llog.FATAL(err)
	}

	defer rows.Close()
	for rows.Next() {
		ev := new(Event)
		if err := rows.Scan(&ev.Id, &ev.AccountId, &ev.Data, &ev.State); err != nil {
			llog.FATAL(err)
		}
		result = append(result, ev)
	}
	if err := rows.Err(); err != nil {
		llog.FATAL(err)
	}
	return result
}

func RegisterEvent(event Event) (Event, error) {
	llog.Success("Attempting 2 Register event")
	registrations := event.Registrations()
	if len(registrations) == 0 {
		return event, errors.New("Nothing registered for this event")
	}
	id, err := db.Create(&event)
	if err != nil {
		return event, errors.New("Nothing registered for this event")
	}
	event.Id = int(id)
	after := time.After(time.Duration(2) * time.Second)
	var state string
	llog.Success("Attempting 2 Q event")
	select {
	case queue <- event:
		llog.Success("Event Queued")
		state = "processing"
	case <-after:
		state = "timeout"
	}
	_, err = db.Exec(
		"UPDATE events SET state = ? WHERE id =	?",
		state,
		id,
	)
	if err != nil {
		llog.Error(err)
		return event, errors.New("Event queued but could not be updated")
	}
	event.State = state
	return event, nil
}
