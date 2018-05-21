package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Pg *sqlx.DB
