package main

import (
	"github.com/awsmsrc/llog"
)

type Attempt struct {
	Id       int    `json:"id,omitempty"`
	Status   int    `json:"status"`
	Response string `json:"reaponse"`
	EventId  int    `json:"event_id,omitempty"`
}

func (a *Attempt) TableName() string {
	return "attempts"
}

func (a *Attempt) Attributes() ([]string, []interface{}) {
	return []string{
			"event_id",
			"response",
			"status",
		}, []interface{}{
			a.EventId,
			a.Response,
			a.Status,
		}
}

func GetAttempts(account_id int) []*Attempt {
	var result []*Attempt
	rows, err := db.Query(
		"SELECT id, event_id, status, response FROM attempts LIMIT 100",
	)
	if err != nil {
		llog.FATAL(err)
	}

	defer rows.Close()
	for rows.Next() {
		at := new(Attempt)
		if err := rows.Scan(&at.Id, &at.EventId, &at.Status, &at.Response); err != nil {
			llog.FATAL(err)
		}
		result = append(result, at)
	}
	return result
}
