package pgclient

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xlzpm/internal/config"
	"github.com/xlzpm/pkg/logger/initlog"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewPostgresDB(ctx context.Context, st config.Storage) (pool *pgxpool.Pool, err error) {
	log := initlog.InitLogger()

	path := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		st.UserName, st.Password, st.Host, st.Port, st.DbName,
	)

	pool, err = pgxpool.New(ctx, path)
	if err != nil {
		log.Error("Don't open pg: ", err)
	}

	return pool, err
}
