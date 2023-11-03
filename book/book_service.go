package book

import (
	"context"

	"github.com/jackc/pgx/v5"

	data "sqlc-example/book/data/gen"
)

type Service struct {
	q *data.Queries
}

func New(conn *pgx.Conn) Service {
	return Service{data.New(conn)}
}

func (s Service) GetBooks(ctx context.Context) ([]data.Book, error) {
	return s.q.ListBooks(ctx)
}
