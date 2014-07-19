package main

import (
	"database/sql"
	"fmt"
	"github.com/awsmsrc/llog"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

type DB struct {
	*sql.DB
}

func (db *DB) Create(obj Attributer) (int64, error) {
	table := obj.TableName()
	keys, values := obj.Attributes()

	tx, err := db.Begin()
	if err != nil {
		llog.Error(err)
		return 0, err
	}

	qs := make([]string, len(values))
	for i, _ := range values {
		qs[i] = "?"
	}

	sql := fmt.Sprintf(
		"insert into %s(%s) values(%s)",
		table,
		strings.Join(keys, ", "),
		strings.Join(qs, ", "),
	)
	llog.Debugf("Generated sql: %v", sql)
	stmt, err := tx.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		llog.Error(err)
		return 0, err
	}

	result, err := stmt.Exec(values...)
	if err != nil {
		llog.Error(err)
		return 0, err
	}

	tx.Commit()

	id, err := result.LastInsertId()
	if err != nil {
		llog.Error(err)
		return 0, err
	}

	return id, nil
}

func (db *DB) Init() {
	var err error
	sqldb, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		llog.FATAL(err)
	}
	db.DB = sqldb
	llog.Success("Ensuring Database Exists")
	err = db.DefaultTable(
		"registrations",
		"account_id INTEGER",
		"key VARCHAR(25)",
		"url TEXT",
	)
	if err != nil {
		llog.FATAL(err)
	}
	err = db.DefaultTable(
		"events",
		"account_id INTEGER",
		"data TEXT",
		"state VARCHAR(50) NOT NULL DEFAULT 'pending'",
	)
	if err != nil {
		llog.FATAL(err)
	}
	err = db.DefaultTable(
		"attempts",
		"event_id INTEGER",
		"response TEXT",
		"status INTEGER",
	)
	if err != nil {
		llog.FATAL(err)
	}
	llog.Success("Database initialized and configured")
}

func (db *DB) DefaultTable(name string, col_defs ...string) error {
	llog.Successf("Creating default table: %v", name)
	_, err := db.Exec(
		fmt.Sprintf(
			"CREATE TABLE IF NOT EXISTS %v(id INTEGER PRIMARY KEY AUTOINCREMENT, %v);",
			name,
			strings.Join(col_defs, ", "),
		),
	)
	if err == nil {
		llog.Successf("Created default table: %v", name)
	}
	return err
}
