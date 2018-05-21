package models

import (
	"errors"
	"log"
	"time"
)

type Log struct {
	ID          int64     `db:"id"`
	Module      string    `db:"module"`
	Action      string    `db:"action"`
	Result      string    `db:"result"`
	Description string    `db:"description"`
	Time        time.Time `db:"time"`
}

func NewLog() *Log {
	l := &Log{}

	return l
}

func (l *Log) Save() error {
	sql := `

		INSERT INTO "logs" (
			module,
			action,
			result,
			description
		)
		VALUES (
			$1,
			$2,
			$3,
			$4
		)

	`
	_, err := Pg.Exec(sql, l.Module, l.Action, l.Result, l.Description)
	if err != nil {
		log.Println("Insert Error: ", err.Error())
		return err
	}

	return nil
}

func (l *Log) Validate() error {
	msg := "Validation Errors:"
	n_errs := 0
	if l.Module == "" {
		msg = msg + " `module` required."
		n_errs = n_errs + 1
	}

	if l.Action == "" {
		msg = msg + " `action` required."
		n_errs = n_errs + 1
	}

	if l.Result == "" {
		msg = msg + " `result` required."
		n_errs = n_errs + 1
	}

	if l.Description == "" {
		msg = msg + " `description` required."
		n_errs = n_errs + 1
	}

	if n_errs > 0 {
		err := errors.New(msg)
		return err
	}
	// else
	return nil
}
