package pg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/xlzpm/internal/users/model"
	"github.com/xlzpm/pkg/logger/initlog"
	"github.com/xlzpm/pkg/pgclient"
)

type Service interface {
	Create(ctx context.Context, user *model.User) error
	FindOne(ctx context.Context, email string, password string) (model.User, error)
}

type Repository struct {
	client pgclient.Client
}

func NewRepository(client pgclient.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *Repository) Create(ctx context.Context, user *model.User) error {
	log := initlog.InitLogger()

	q := `	INSERT INTO users
						(email, password)
			VALUES 
						($1, $2)
			RETURNING id`

	log.Info(fmt.Sprintf("SQL query: %s", formatQuery(q)))

	if err := r.client.QueryRow(ctx, q, user.Email, user.Password).Scan(&user.Id); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Error(newErr.Error())
			return newErr
		}

		return nil
	}

	return nil
}

func (r *Repository) FindOne(ctx context.Context, email string, password string) (model.User, error) {
	log := initlog.InitLogger()

	q := `SELECT email, password FROM users WHERE email = $1 AND password = $2`

	log.Info(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var user model.User
	err := r.client.QueryRow(ctx, q, email, password).Scan(&user.Email, &user.Password)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
