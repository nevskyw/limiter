package database

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
)

type Conn interface {
	Exec(sql squirrel.Sqlizer, arguments ...interface{}) (err error)
}

type conn struct {
	pg *pgx.Conn
}

func newConn(c *pgx.Conn) Conn {
	return &conn{
		pg: c,
	}
}

func (c *conn) Exec(sql squirrel.Sqlizer, arguments ...interface{}) (err error) {
	q, args, err := sql.ToSql()
	if err != nil {
		return err
	}

	_, err = c.pg.Exec(q, args)

	return err
}
