package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xlzpm/internal/config"
	"github.com/xlzpm/pkg/logger/initlog"
)

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
