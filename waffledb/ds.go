package waffledb

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func getPGConn(ctx context.Context) *pgx.Conn {
	// TODO: extract connection string to be sourced from env
	config, err := pgx.ParseConfig("postgresql://postgres@localhost:5432/waffle")
	if err != nil {
		panic("parse config failure")
	}
	conn, err := pgx.ConnectConfig(ctx, config)
	if err != nil {
		panic("failed to connect")
	}
	return conn
}

type WaffleDB struct {
	conn *pgx.Conn
}

func MustGetDS(ctx context.Context) WaffleDB {
	return WaffleDB{getPGConn(ctx)}
}
