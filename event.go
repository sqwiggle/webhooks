package main

import (
	"github.com/awsmsrc/llog"
	"time"
	"errors"
)

type Event struct {
	Id        int       `json:"id"`
	AccountId int       `json:"account_id"`
	Data      string    `json:"data"`
	Attempts  *[]*Event `json:"events,omitempty"`
	State     string   `json:"state"`
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

func (e *Event) InterestedUrls() map[int]string {
	result := make(map[int]string)
	rows, err := db.Query(
		"SELECT id, url FROM registrations WHERE account_id=?",
		e.AccountId,
	)
	if err != nil {
		llog.FATAL(err)
	}

	defer rows.Close()
	for rows.Next() {
		var url string
		var id int
		if err := rows.Scan(&id, &url); err != nil {
			llog.FATAL(err)
		}
		result[id] = url
	}
	if err := rows.Err(); err != nil {
		llog.FATAL(err)
	}
	return result
}

func GetEvents (account_id int) []*Event {
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

func RegisterEvent (event Event) (Event, error) {
	llog.Success("Attempting 2 Register event")
	urls := event.InterestedUrls()
	if len(urls) == 0 {
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
	case <- after:
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
