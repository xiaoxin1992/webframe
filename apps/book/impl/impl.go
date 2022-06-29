package impl

import "database/sql"

var (
	svr = &service{}
)

type service struct {
	db *sql.DB
}
