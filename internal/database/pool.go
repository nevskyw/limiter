package database

import (
	"context"
	"github.com/jackc/pgx"
	"limiter/internal/tools/closer"
	"time"
)

type Pool interface {
	Acquire() (Conn, error)
}

type pool struct {
	p *pgx.ConnPool
}

// New returns new pool
func New(ctx context.Context) (Pool, error) {
	conn, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "",
			Port:     0,
			Database: "",
			User:     "",
			Password: "",
		},
		MaxConnections: 0,
		AcquireTimeout: 0,
	})
	if err != nil {
		return nil, err
	}

	ctxT, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	c, err := conn.Acquire()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = c.Close()
	}()

	p := &pool{conn}

	closer.Closer.Add(p.close)

	return p, c.Ping(ctxT)
}

func (p *pool) Acquire() (Conn, error) {
	c, err := p.p.Acquire()
	if err != nil {
		return nil, err
	}
	return newConn(c), nil
}

func (p *pool) close() error {
	p.p.Close()
	return nil
}
