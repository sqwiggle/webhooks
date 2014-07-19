package main

import (
	"github.com/awsmsrc/llog"
)

type Registration struct {
	Id        int    `json:"id"`
	Key        string    `json:"key"`
	AccountId int    `json:"account_id"`
	Url       string `json:"url"`
}

func (r *Registration) TableName() string {
	return "registrations"
}

func (r *Registration) Attributes() ([]string, []interface{}) {
	return []string{
			"account_id",
			"key",
			"url",
		}, []interface{}{
			r.AccountId,
			r.Key,
			r.Url,
		}
}

func GetRegistrations (account_id int) []*Registration {
	var result []*Registration
	rows, err := db.Query(
		"SELECT id, account_id, url FROM registrations WHERE account_id=? LIMIT 100",
		account_id,
	)
	if err != nil {
		llog.FATAL(err)
	}

	defer rows.Close()
	for rows.Next() {
		re := new(Registration)
		if err := rows.Scan(&re.Id, &re.AccountId, &re.Url); err != nil {
			llog.FATAL(err)
		}
		result = append(result, re)
	}
	return result
}
